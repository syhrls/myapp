package handlers

import (
	"example/hello/utils"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	utils.InfoWithContext(c, "Fetching dummy user data")

	// Dummy data
	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Smith"},
	}

	utils.SuccessResponse(c, "Users fetched successfully", users)
}