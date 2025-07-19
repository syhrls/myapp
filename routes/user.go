package routes

import (
	"errors"
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			err := errors.New("Bad Request")
			utils.LogError(err)
			utils.BadRequestResponse(c, err.Error())
		})
	}
}
