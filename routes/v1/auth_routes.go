package v1

import (
	"example/hello/database"
	"example/hello/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/v1/auth")
	{
		auth.POST("/register", handlers.RegisterHandler(database.DB))
		auth.POST("/login", handlers.LoginHandler(database.DB))
	}
}
