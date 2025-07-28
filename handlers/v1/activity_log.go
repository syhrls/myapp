package handlers

import (
    "time"

    "example/hello/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// LogActivity mencatat aktivitas user yang sudah login
func LogActivity(db *gorm.DB, userID uuid.UUID, device, action string) {
    // Hapus log yang lebih dari 1 hari
    db.Where("created_at < ?", time.Now().Add(-24*time.Hour)).Delete(&models.ActivityLog{})

    // Simpan log baru
    log := models.ActivityLog{
        ID:        uuid.New(),
        UserID:    userID,
        Device:    device,
        Action:    action,
        CreatedAt: time.Now(),
    }
    db.Create(&log)
}