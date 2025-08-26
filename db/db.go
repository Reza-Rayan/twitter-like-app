package db

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/config"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.AppConfig.Database

	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.Name), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to db: %v", err))
	}

	// Create Tables Based on Models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Follow{},
		&models.Notification{},
		&models.Like{},
		&models.OTP{},
		&models.Role{},
	)
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}
}
