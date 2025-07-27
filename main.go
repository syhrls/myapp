package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example/hello/database"
	"example/hello/middleware"
	"example/hello/routes"
	"example/hello/utils"
)

func main() {
	// Inisialisasi logger
	utils.InitLoggerWIB()
	gin.DefaultWriter = utils.ColorWriter{Writer: os.Stdout}

	// Load environment
	if err := godotenv.Load(); err != nil {
		utils.Error(".env file not found or failed to load")
	} else {
		utils.Info(".env loaded successfully")
	}

	// Inisialisasi database
	database.InitMySQL()

	// Register UUID hook setelah DB siap
	utils.RegisterGlobalUUIDHook(database.DB)

	// Migrasi database
	database.MigrateModels(database.DB)

	// Set Gin mode
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	utils.Info("Running in mode: " + mode)

	// Setup Router dan Middleware
	r := gin.Default()
	r.Use(middleware.RequestID())

	// Setup Trusted Proxies dan Routes
	r.SetTrustedProxies([]string{"192.168.1.2"})
	routes.SetupRoutes(r)

	// Handler untuk route yang tidak ditemukan
	r.NoRoute(func(c *gin.Context) {
		utils.ErrorResponse(c, utils.CodeNotFound, "No Route Matched")
	})

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	utils.Info("Server running on port: " + port)
	r.Run(":" + port)
}
