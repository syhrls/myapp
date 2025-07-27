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

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.
		Where("username = ?", username).
		Order("created_at DESC").
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
