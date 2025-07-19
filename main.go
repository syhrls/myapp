package main

import (
	"errors"
	"example/hello/routes"
	"example/hello/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load")
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.Default()    // gunakan gin.New() agar tidak include logger default

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
