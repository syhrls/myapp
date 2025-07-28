package v1

import (
	"example/hello/database"
	"example/hello/handlers/v1"
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/v1/auth", utils.LogStartEnd())
	{
		auth.POST("/register", handlers.RegisterHandler(database.DB))
		auth.POST("/login", handlers.LoginHandler(database.DB))
	}
}
