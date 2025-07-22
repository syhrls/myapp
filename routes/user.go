package routes

import (
	"example/hello/handlers"
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/users", LogStartEnd())
	{
		user.GET("/", handlers.GetAllUsers)
	}
}

func LogStartEnd() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.InfoWithContext(c, "Start processing %s", c.Request.URL.Path)
		c.Next()
		utils.InfoWithContext(c, "End processing %s", c.Request.URL.Path)
	}
}
