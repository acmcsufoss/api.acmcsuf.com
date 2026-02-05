package middleware

import (
	"net/http"
	"sync"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	maxRate  = 1
	maxBurst = 5
)

var (
	clients = make(map[string]*rate.Limiter)
	mu      sync.Mutex
)

// The rate limiter is an important middleware that
// limits how many times a client can access our server
// per second. This is useful for preventing spam that
// overloads our server
func Ratelimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if config.Load().Env == "development" {
			ctx.Next()
			return
		}

		ip := ctx.ClientIP()
		rl := getClient(ip)

		if rl.Allow() {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests from client",
			})
		}
	}
}

// Each client is able to send 5 requests a burst
// While also gaining 1 request per second
func getClient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if lim, ok := clients[ip]; ok {
		return lim
	}

	lim := rate.NewLimiter(maxRate, maxBurst)
	clients[ip] = lim

	return lim
}
