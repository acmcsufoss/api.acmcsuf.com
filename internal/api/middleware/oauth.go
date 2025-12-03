package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

var RoleMap = map[string]string{
	"123": "Board",
	"456": "President",
}

var roleCache = sync.Map{}

type cacheEntry struct {
	Roles     []string
	UserID    string
	ExpiresAt time.Time
}

func DiscordAuth(bot *discordgo.Session, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// expects the header 'Authorization: Bearer <access_token>'
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			return
		}

		// dev mode bypass (ENV=development)
		cfg := config.Load()
		if cfg.Env == "development" && authHeader == "Bearer dev-token" {
			c.Set("userID", "dev-user-id")
			c.Next()
			return
		}

		// using a cache is required since discord has pretty strict rate limits on their API
		if value, ok := roleCache.Load(authHeader); ok {
			cached := value.(cacheEntry)
			if time.Now().Before(cached.ExpiresAt) {
				if checkRoles(cached.Roles, requiredRole) {
					c.Set("userID", cached.UserID)
					c.Next()
					return
				} else {
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
						"error": "Insufficient permissions (cached)",
					})
					return
				}
			} else {
				roleCache.Delete(authHeader)
			}

		}

		userSession, _ := discordgo.New(authHeader)
		user, err := userSession.User("@me")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired Discord token",
			})
			return
		}

		member, err := bot.GuildMember(cfg.GuildID, user.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are not a member of the Discord server",
			})
			return
		}

		roleCache.Store(authHeader, cacheEntry{
			Roles:     member.Roles,
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(time.Minute * 5),
		})

		if checkRoles(member.Roles, requiredRole) {
			c.Set("userID", user.ID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Insufficient permissions",
			})
		}

	}
}

func checkRoles(userRoleIDs []string, requiredRole string) bool {
	for _, id := range userRoleIDs {
		roleName, exists := RoleMap[id]
		if !exists {
			continue
		}

		if roleName == requiredRole {
			return true
		}
		if roleName == "President" && requiredRole == "Board" {
			return true
		}
	}
	return false
}
