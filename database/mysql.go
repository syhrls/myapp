package database

import (
	"fmt"
	"os"

	"example/hello/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	utils.Debug("DSN: " + dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Fatal("Failed to connect to MySQL: " + err.Error())
	}

	if DB == nil {
		utils.Fatal("GORM returned nil DB instance")
	}

	utils.Info("Connected to MySQL")
}
