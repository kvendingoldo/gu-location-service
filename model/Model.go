package model

import (
	"github.com/kvendingoldo/gu-location-service/config"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int64     `json:"id" example:"23" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

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
