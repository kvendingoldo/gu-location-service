package model

import "time"

type Location struct {
	ID          int64  `json:"id" example:"23" gorm:"primaryKey"`
	UID         int64  `json:"uid" example:"12"` // gorm:foreignKey
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"`

	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
