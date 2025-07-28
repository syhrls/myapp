package handlers

import (
	"example/hello/database"
	"example/hello/dto/v1"
	"example/hello/models"
	"example/hello/repositories/v1"
	"example/hello/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllUsers(c *gin.Context) {
	users, err := v1.GetAllUsers()
	if err != nil {
		utils.BadRequestResponse(c, "Failed to retrieve users: "+err.Error())
		return
	}

	if len(users) == 0 {
		utils.SuccessResponse(c, "No users found", nil)
		return
	}

	filteredUsers := make([]map[string]any, len(users))
	for i, u := range users {
		filteredUsers[i] = map[string]any{
			"username": u.Username,
			"email":    u.Email,
		}
	}

	userID, _ := c.Get("user_id")
	device := c.GetHeader("X-Device-Id")
	if device == "" {
		device = c.GetHeader("User-Agent") // fallback jika tidak ada
	}
	LogActivity(database.DB, userID.(uuid.UUID), device, c.Request.Method+" "+c.Request.URL.Path)

	utils.SuccessResponse(c, "Users retrieved successfully", filteredUsers)
}

func CreateUser(c *gin.Context) {
	var input dto.CreateUserDTO

	// Validasi request body
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid input data")
		return
	}

	if condition := input.Username == ""; condition {
		utils.BadRequestResponse(c, "Username is required")
		return
	}

	// Map DTO ke model User
	user := models.User{
		Username:  input.Username,
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

	// Only return username and email fields
	filteredUser := map[string]any{
		"username": user.Username,
		"email":    user.Email,
	}

	userID, _ := c.Get("user_id")
device := c.GetHeader("X-Device-Id")
	if device == "" {
		device = c.GetHeader("User-Agent") // fallback jika tidak ada
	}	
	LogActivity(database.DB, userID.(uuid.UUID), device, c.Request.Method+" "+c.Request.URL.Path)

	utils.SuccessResponse(c, "User found", filteredUser)

}
