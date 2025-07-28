package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"example/hello/models"
	v1 "example/hello/repositories/v1"
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
			Username:  req.Username,
			CreatedBy: utils.SYSTEM,
			UpdatedBy: utils.SYSTEM,
			Password:  hash,
			Salt:      salt,
		}
		if err := db.Create(&user).Error; err != nil {
			utils.BadRequestResponse(c, "Failed to register: "+err.Error())
			return
		}
		utils.SuccessResponse(c, "Registration successful", nil)
	}
}

// GenerateRandomToken membuat random string 16 digit (32 karakter hex)
func GenerateRandomToken() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// LoginHandler untuk login dan generate token random, simpan ke UserToken
func LoginHandler(db *gorm.DB, userRepo v1.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Invalid input data")
			return
		}

		user, err := userRepo.FindByUsername(req.Username)
		if err != nil {
			utils.BadRequestResponse(c, "Username not found")
			return
		}

		if !utils.CheckPassword(user.Password, req.Password, user.Salt) {
			utils.BadRequestResponse(c, "Invalid password")
			return
		}

		token, err := GenerateRandomToken()
		if err != nil {
			utils.BadRequestResponse(c, "Failed to generate token")
			return
		}
		refreshToken, err := GenerateRandomToken()
		if err != nil {
			utils.BadRequestResponse(c, "Failed to generate refresh token")
			return
		}

		expired := time.Now().Add(24 * time.Hour).Unix()

		userToken := models.UserToken{
			UserID:    user.ID,
			Token:     token,
			ExpiredAt: time.Unix(expired, 0),
			CreatedAt: time.Now(),
		}
		if err := db.Create(&userToken).Error; err != nil {
			utils.BadRequestResponse(c, "Failed to save token")
			return
		}

		utils.SuccessResponse(c, "Login successful", gin.H{
			"expired":       expired,
			"refresh_token": refreshToken,
			"token":         token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		})
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