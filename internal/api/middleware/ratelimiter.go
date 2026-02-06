package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	maxRate  = 100
	maxBurst = 200
)

type client struct {
	rl  *rate.Limiter
	ttl time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
	once    sync.Once
)

// The rate limiter is an important middleware that
// limits how many times a client can access our server
// per second. This is useful for preventing spam that
// overloads our server
func Ratelimiter() gin.HandlerFunc {
	once.Do(func() {
		go checkTTL()
	})
	environment := config.Load().Env
	return func(ctx *gin.Context) {
		if environment == "development" {
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

// Since we don't want to run out of memory while storing our clients
// We will keep track of our clients for 24 hours, if they have not requested
// in 24 hours, we will remove them from the table
func checkTTL() {
	for {
		time.Sleep(5 * time.Minute)
		now := time.Now()
		mu.Lock()
		for key, elm := range clients {
			if elm.ttl.Before(now) {
				delete(clients, key)
			}
		}
		mu.Unlock()
	}
}

// Each client is able to send 200 requests a burst
// While also gaining 100 request per second
func getClient(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if lim, ok := clients[ip]; ok {
		newTTL := time.Now().Add(2 * time.Hour)
		lim.ttl = newTTL
		return lim.rl
	}

	lim := rate.NewLimiter(maxRate, maxBurst)
	newTTL := time.Now().Add(2 * time.Hour)

	newClient := &client{
		rl:  lim,
		ttl: newTTL,
	}

	clients[ip] = newClient

	return lim
}
