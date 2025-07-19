package main

import (
	"errors"
	"example/hello/routes"
	"example/hello/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()                    // gunakan gin.New() agar tidak include logger default
	r.Use(gin.Recovery())             // tangani panic

	routes.SetupRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		err := errors.New("route not found")
		utils.LogError(err)
		utils.ErrorResponse(c, utils.CodeNotFound, "Route not found")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)
	r.Run(":" + port)
}
