package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type endpointVisitor struct {
	Requests int
	ResetAt  time.Time
}

var endpointVisitors = make(map[string]map[string]*endpointVisitor) // map[ip][endpoint]
var mu sync.Mutex

func RateLimitPerEndpoint(maxReq int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		endpoint := c.FullPath() // endpoint unik per handler
		now := time.Now()

		mu.Lock()
		if endpointVisitors[ip] == nil {
			endpointVisitors[ip] = make(map[string]*endpointVisitor)
		}
		v, exists := endpointVisitors[ip][endpoint]
		if !exists || now.After(v.ResetAt) {
			endpointVisitors[ip][endpoint] = &endpointVisitor{Requests: 1, ResetAt: now.Add(window)}
			mu.Unlock()
			c.Next()
			return
		}

		if v.Requests >= maxReq {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded for this endpoint. Try again in 1 minute.",
			})
			return
		}
		v.Requests++
		mu.Unlock()
		c.Next()
	}
}
