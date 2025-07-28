package routes

import (
	"example/hello/middleware"
	"example/hello/routes/v1"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    api.Use(middleware.RateLimit(5, time.Minute)) // Rate limit: 5x per menit per IP
    v1.SetupUserRoutes(api)
    v1.SetupAuthRoutes(api)
}