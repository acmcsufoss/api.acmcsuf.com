package middleware

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS is a way to provide security to our API, it decides what sources
// can access our routes, what HTTP requests they could send, and much more.
func Cors() gin.HandlerFunc {
	cfg := config.Load()

	newCors := cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// Let us discuss the following fields
		AllowHeaders:     []string{},
		ExposeHeaders:    []string{},
		AllowCredentials: false,
		MaxAge:           time.Hour * 24,
	})

	return newCors
}
