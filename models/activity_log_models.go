package models

import (
    "time"

    "github.com/google/uuid"
)

type ActivityLog struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
    Device    string    `gorm:"type:varchar(255)" json:"device"`
    Action    string    `gorm:"type:varchar(255)" json:"action"`
    CreatedAt time.Time `json:"created_at"`

    User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}