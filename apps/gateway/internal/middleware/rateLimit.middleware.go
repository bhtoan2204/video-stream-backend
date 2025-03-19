package middleware

import (
	"sync"
	"time"

	"github.com/bhtoan2204/gateway/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type rateLimiter struct {
	clients map[string]*client
	mu      sync.Mutex

	requestsPerSecond rate.Limit
	burstSize         int
}

func NewRateLimiter(r rate.Limit, b int) *rateLimiter {
	rl := &rateLimiter{
		clients:           make(map[string]*client),
		requestsPerSecond: r,
		burstSize:         b,
	}

	go rl.cleanupClients()

	return rl
}

func (rl *rateLimiter) getClient(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if c, exists := rl.clients[ip]; exists {
		c.lastSeen = time.Now()
		return c.limiter
	}

	limiter := rate.NewLimiter(rl.requestsPerSecond, rl.burstSize)
	rl.clients[ip] = &client{
		limiter:  limiter,
		lastSeen: time.Now(),
	}
	return limiter
}

func (rl *rateLimiter) cleanupClients() {
	for {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		for ip, c := range rl.clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}
func RateLimitMiddleware(rl *rateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rl.getClient(ip)

		if !limiter.Allow() {
			response.ErrorTooManyRequestsResponse(c, 429)
			c.Abort()
			return
		}

		c.Next()
	}
}
