package v1

import (
	"example/hello/database"
	"example/hello/models"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// FindByUsername mencari user berdasarkan username (exact match)
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
