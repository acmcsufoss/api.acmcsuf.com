package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
)

// Cors is a way to provide security to our API
// Using cors we can say what we want to access our api
// Like from specific sites for example.
func Cors() gin.HandlerFunc {

	cfg := config.Load()

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if !allowedOrigin(origin, cfg) {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("request is not allowed from this origin: %s", origin))
			return
		}

		corsCfg := cors.New(cors.Config{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
			AllowHeaders:     []string{},
			ExposeHeaders:    []string{},
			AllowCredentials: false,
			MaxAge:           time.Hour * 24,
		})

		corsCfg(c)

		c.Next()
	}
}

// Later when deployed, this can block dev origin and only accept ones used in prod
func allowedOrigin(o string, cfg *config.Config) bool {
	if cfg.Env == "production" && o == "development" {
		return false
	}

	if slices.Contains(cfg.AllowedOrigins, o) {
		return true
	}
	return false
}
