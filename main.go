package main

import (
	"example/hello/routes"
	"example/hello/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	// Handler untuk route yang tidak ditemukan
	r.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, utils.CodeNotFound, "Route not found")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Println("Server running on port:", port)
	r.Run(":" + port)
}
