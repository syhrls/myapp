package v1

import (
	"example/hello/database"
	"example/hello/handlers/v1"
	"example/hello/middleware"
	"example/hello/utils"
	"time"

	"github.com/gin-gonic/gin"
)
func SetupUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/v1/users", utils.LogStartEnd(), middleware.TokenAuthMiddleware(database.DB))
	{
		user.GET("/", middleware.RateLimitPerEndpoint(5, time.Minute), handlers.GetAllUsers)
		user.POST("/", middleware.RateLimitPerEndpoint(5, time.Minute), handlers.CreateUser)
		user.GET("/search", middleware.RateLimitPerEndpoint(5, time.Minute), handlers.GetUserByUsername)
	}
}
