package middleware

import (
    "os"
    "strings"

    "example/hello/utils"
    "github.com/gin-gonic/gin"
)

// JWTAuthMiddleware memvalidasi token JWT pada header Authorization
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            utils.ErrorResponse(c, utils.CodeUnauthorized, "Missing or invalid Authorization header")
            c.Abort()
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        secret := os.Getenv("JWT_SECRET")
        claims, err := utils.ParseJWT(tokenString, secret)
        if err != nil {
            utils.ErrorResponse(c, utils.CodeUnauthorized, "Invalid or expired token")
            c.Abort()
            return
        }
        // Simpan claims ke context jika perlu
        c.Set("user", claims)
        c.Next()
    }
}