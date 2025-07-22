package routes

import (
	"example/hello/handlers"
	"example/hello/utils"
	"fmt"

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

		utils.InfoWithContext(c, "%s", fmt.Sprintf("Start processing %s", c.FullPath()))

		c.Next()

		utils.InfoWithContext(c, "%s", fmt.Sprintf("End processing %s", c.FullPath()))
	}
}