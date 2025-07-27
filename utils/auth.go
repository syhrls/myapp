package utils

import (
	"maps"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// GenerateSalt creates a random salt string
func GenerateSalt(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// HashPassword hashes the password with salt using bcrypt
func HashPassword(password, salt string) (string, error) {
	salted := password + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword compares hash with password+salt
func CheckPassword(hash, password, salt string) bool {
	salted := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salted))
	return err == nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(claims map[string]any, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	mapClaims := jwt.MapClaims{}
	maps.Copy(mapClaims, claims)
	token.Claims = mapClaims
	return token.SignedString([]byte(secret))
}

// ParseJWT parses and validates JWT token
func ParseJWT(tokenStr, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid claims")
}

func GetJWTExpiration() int64 {
	// Set token to expire in 3 hours
	return time.Now().Add(3 * time.Hour).Unix()
}
