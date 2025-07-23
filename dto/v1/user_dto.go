package dto

type CreateUserDTO struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	CreatedBy string `json:"created_by" binding:"required"`
}
