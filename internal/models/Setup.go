package models

import (
	"github.com/kvendingoldo/gu-location-service/config"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {

	db = config.Config.DB

	if err := db.AutoMigrate(&Location{}); err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
