package handlers

import (
	"os"

	"example/hello/models"
	"example/hello/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// DTO untuk login
type LoginRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// DTO untuk register
type RegisterRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterHandler untuk daftar akun baru
func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Invalid input data")
			return
		}

		// Cek apakah username sudah ada
		test, err := userRepo.FindByUsername(req.Username)
		if err != nil && err != gorm.ErrRecordNotFound {
			utils.BadRequestResponse(c, "Failed to check username: "+err.Error())
			return
		}

		if test != nil {
			utils.BadRequestResponse(c, "Username already exists")
			return
		}

		salt, _ := utils.GenerateSalt(16)
		hash, _ := utils.HashPassword(req.Password, salt)

		user := models.User{
			Username: req.Username,
			CreatedBy: utils.SYSTEM,
			UpdatedBy: utils.SYSTEM,
			Password: hash,
			Salt:     salt,
		}
		if err := db.Create(&user).Error; err != nil {
			utils.BadRequestResponse(c, "Failed to register: "+err.Error())
			return
		}
		utils.SuccessResponse(c, "Registration successful", nil)
	}
}

// LoginHandler untuk login dan mendapatkan JWT
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Invalid input data")
			return
		}

		// Ambil user dari repository
		user, err := userRepo.FindByUsername(req.Username)
		if err != nil {
			utils.UnauthorizedResponse(c, "Username not found")
			return
		}

		if !utils.CheckPassword(user.Password, req.Password, user.Salt) {
			utils.UnauthorizedResponse(c, "Invalid password")
			return
		}

		secret := os.Getenv("JWT_SECRET")
		exp := utils.GetJWTExpiration()
		claims := map[string]any{
			"id":       user.ID.String(),
			"username": user.Username,
			"email":    user.Email,
			"exp":      exp,
		}
		token, err := utils.GenerateJWT(claims, secret)
		if err != nil {
			utils.BadRequestResponse(c, "Failed to generate token: "+err.Error())
			return
		}

		// Generate refresh token (could be another JWT or random string)
		refreshExp := utils.GetJWTExpiration() / 60
		refreshClaims := map[string]any{
			"id":       user.ID.String(),
			"username": user.Username,
			"email":    user.Email,
			"exp":      refreshExp,
			"type":     "refresh",
		}
		refreshToken, err := utils.GenerateJWT(refreshClaims, secret)
		if err != nil {
			utils.BadRequestResponse(c, "Failed to generate refresh token: "+err.Error())
			return
		}

		response := gin.H{
			"user": gin.H{
				"id":       user.ID.String(),
				"username": user.Username,
				"email":    user.Email,
			},
			"token":         token,
			"refresh_token": refreshToken,
			"expired":       exp,
		}
		utils.SuccessResponse(c, "Login successful", response)
	}
}

func RefreshTokenHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Token string `json:"token" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Invalid input data")
			return
		}

		secret := os.Getenv("JWT_SECRET")
		claims, err := utils.ParseJWT(req.Token, secret)
		if err != nil {
			utils.UnauthorizedResponse(c, "Invalid or expired token")
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			utils.BadRequestResponse(c, "Invalid token claims")
			return
		}

		user, err := userRepo.FindByUsername(username)
		if err != nil {
			utils.UnauthorizedResponse(c, "User not found")
			return
		}

		newClaims := map[string]any{
			"id":       user.ID.String(),
			"username": user.Username,
			"email":    user.Email,
			"exp":      utils.GetJWTExpiration(),
		}
		newToken, err := utils.GenerateJWT(newClaims, secret)
		if err != nil {
			utils.BadRequestResponse(c, "Failed to generate token: "+err.Error())
			return
		}

		response := gin.H{
			"user": gin.H{
				"id":       user.ID.String(),
				"username": user.Username,
				"email":    user.Email,
			},
			"token": newToken,
		}
		utils.SuccessResponse(c, "Token refreshed successfully", response)
	}
}