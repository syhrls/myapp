package middleware

import (
	"strings"
	"time"

	"example/hello/models"
	"example/hello/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TokenAuthMiddleware memvalidasi token dari header Authorization dengan DB
func TokenAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.UnauthorizedResponse(c, "Missing or invalid Authorization header")
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		var userToken models.UserToken
		if err := db.Where("token = ? AND is_revoked = ?", token, false).First(&userToken).Error; err != nil {
			utils.UnauthorizedResponse(c, "Invalid or revoked token")
			return
		}
		if userToken.ExpiredAt.Before(time.Now()) {
			utils.UnauthorizedResponse(c, "Token expired")
			return
		}
		// Optionally, set user info to context
		c.Set("user_id", userToken.UserID)
		c.Next()
	}
}
