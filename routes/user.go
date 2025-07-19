package routes

import (
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			utils.BadRequestResponse(c, utils.CodeBadRequest, "Bad Request")
		})
	}
}
