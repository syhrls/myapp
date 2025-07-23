package handlers

import (
	"example/hello/database"
	"example/hello/dto/v1"
	"example/hello/models"
	"example/hello/repositories/v1"
	"example/hello/utils"
	"time"

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

func CreateUser(c *gin.Context) {
	var input dto.CreateUserDTO

	// Validasi request body
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input data")
		return
	}

	if condition := input.Username == "" || input.Email == ""; condition {
		utils.BadRequestResponse(c, "Username and Email are required")
		return
	}

	// Map DTO ke model User
	user := models.User{
		Username:  input.Username,
		Email:     input.Email,
		CreatedBy: utils.SYSTEM_CAPS,
		CreatedAt: time.Now(),
		UpdatedBy: utils.SYSTEM_CAPS,
		UpdatedAt: time.Now(),
	}

	// Simpan ke database
	if err := database.DB.Create(&user).Error; err != nil {
		utils.BadRequestResponse(c, "Failed to create user: "+err.Error())
		return
	}

	utils.SuccessResponse(c, "User created successfully", user)
}

var userRepo = v1.NewUserRepository()

func GetUserByUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		utils.BadRequestResponse(c, "Username is required")
		return
	}

	user, err := userRepo.FindByUsername(username)
	if err != nil {
		utils.BadRequestResponse(c, "User not found: "+err.Error())
		return
	}

	utils.SuccessResponse(c, "User found", user)
}
