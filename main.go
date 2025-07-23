package main

import (
	"example/hello/database"
	"example/hello/middleware"
	"example/hello/models"
	"example/hello/routes"
	"example/hello/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	utils.InitLoggerWIB()
	gin.DefaultWriter = utils.ColorWriter{Writer: os.Stdout}

	// Load .env
	err := godotenv.Load()
	if err != nil {
		utils.Error(".env file not found or failed to load")
	} else {
		utils.Info(".env loaded successfully")
	}

	// Init DB
	database.InitMySQL()

	// DB Migration
	if database.DB != nil {
		err := database.DB.AutoMigrate(&models.User{})
		if err != nil {
			utils.Fatal("Auto migration failed: " + err.Error())
		}
		utils.Info("Database migrated successfully")
	}

	// Set Gin mode
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}

	// Set Gin mode
	gin.SetMode(mode)
	utils.Info("Running in mode: " + mode)

	// Setup Router
	r := gin.Default()
	r.Use(middleware.RequestID())

	// Setup Routes
	r.SetTrustedProxies([]string{"192.168.1.2"})
	routes.SetupRoutes(r)

	// NoRoute fallback
	r.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, utils.CodeNotFound, "No Route Matched")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	utils.Info("Server running on port: " + port)
	r.Run(":" + port)
}
