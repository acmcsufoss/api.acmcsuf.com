package middleware

import (
	"fmt"
	"net/http"
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
	fmt.Println("RUN CORS----------------------------------------------------\n", cfg.AllowedOrigins)

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> ORIGIN:", origin)

		if !allowedOrigin(origin, cfg) {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("request is not allowed from this origin"))
			return
		}

		cors.New(cors.Config{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
			AllowHeaders:     []string{},
			ExposeHeaders:    []string{},
			AllowCredentials: false,
			MaxAge:           time.Hour * 24,
		})

		c.Next()
	}
}

func allowedOrigin(o string, cfg *config.Config) bool {
	for _, elm := range cfg.AllowedOrigins {
		if elm == o {
			return true
		}
	}
	return false
}
