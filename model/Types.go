package model

type Location struct {
	Model
	UID         int64  `json:"uid" example:"12"` // gorm:foreignKey
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
