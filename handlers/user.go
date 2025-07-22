package handlers

import (
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	utils.Info("Fetching all users")

	// Dummy data
	users := []map[string]any{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Smith"},
	}

	utils.SuccessResponse(c, "Users fetched successfully", users)
}