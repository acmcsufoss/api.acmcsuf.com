package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

var roleCache = sync.Map{}

type cacheEntry struct {
	Roles     []string
	UserID    string
	expiresAt time.Time
}

func DiscordAuthMiddleware(bot *discordgo.Session, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// expects the header 'Authorization: Bearer <access_token>'
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// dev mode bypass (ENV=development)
		if config.Load().Env == "development" && authHeader == "Bearer dev-token" {
			c.Set("userID", "dev-user-id")
			c.Next()
			return
		}
	}

}
