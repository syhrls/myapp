package main

import (
	"example/hello/routes"
	"example/hello/utils"
	"example/hello/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	r := gin.New() // gunakan gin.New() agar tidak include logger default
	r.Use(gin.Recovery())          // tangani panic
	r.Use(middleware.RequestLogger()) // logger kustom kamu

	routes.SetupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		utils.Logger.Warn("Route not found: " + c.Request.URL.Path)
		utils.ErrorResponse(c, utils.CodeNotFound, "Route not found")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)
	r.Run(":" + port)
}
