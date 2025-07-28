package middleware

import (
    "net/http"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
)

type visitor struct {
    Requests int
    ResetAt  time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

func RateLimit(maxReq int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        now := time.Now()

        mu.Lock()
        v, exists := visitors[ip]
        if !exists || now.After(v.ResetAt) {
            visitors[ip] = &visitor{Requests: 1, ResetAt: now.Add(window)}
            mu.Unlock()
            c.Next()
            return
        }

        if v.Requests >= maxReq {
            mu.Unlock()
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
                "error": "Rate limit exceeded. Try again in 1 minute.",
            })
            return
        }
        v.Requests++
        mu.Unlock()
        c.Next()
    }
}