package models

import (
	"time"

	"github.com/google/uuid"
)

// UserToken menyimpan token login user (misal untuk session atau refresh token)
type UserToken struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user"` // relasi ke User
}
