package v1

import (
	"example/hello/handlers"
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/v1/users", utils.LogStartEnd())
	{
		user.GET("/", handlers.GetAllUsers)
		user.POST("/", handlers.CreateUser)
		user.GET("/search", handlers.GetUserByUsername)
	}
}
