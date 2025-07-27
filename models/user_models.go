package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"` // gunakan char(36) untuk MySQL
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Email     string    `json:"email"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}