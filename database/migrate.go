package database

import (
	"example/hello/models"
	"example/hello/utils"

	"gorm.io/gorm"
)

// MigrateModels migrates all models if not already migrated
func MigrateModels(db *gorm.DB) {
	modelsToMigrate := []any{
		&models.User{},
		&models.UserToken{},
		&models.ActivityLog{},
		// Tambahkan model lain di sini jika ada, contoh:
		// &models.Product{},
	}

	for _, model := range modelsToMigrate {
		if !db.Migrator().HasTable(model) {
			err := db.AutoMigrate(model)
			if err != nil {
				utils.Fatal("Auto migration failed: " + err.Error())
			}
			utils.Info("Migrated: " + db.Migrator().CurrentDatabase())
		} else {
			utils.Info("Table already migrated, skipping: " + db.Migrator().CurrentDatabase())
		}
	}
}
