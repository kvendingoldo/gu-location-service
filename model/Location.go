package model

import (
	"encoding/json"
	"github.com/kvendingoldo/gu-location-service/config"
	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"

	"time"
)

type Location struct {
	ID          int64  `json:"id" example:"23" gorm:"primaryKey"`
	UID         int64  `json:"uid" example:"12"` // gorm:foreignKey
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"`

	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
}

// UpdateLocation ... Update location
func UpdateLocation(location *Location) (err error) {
	err = config.Config.DB.Create(location).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError guErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return err
		}

		switch newError.Number {
		case 1062:
			err = guErrors.NewAppErrorWithType(guErrors.ResourceAlreadyExists)
			return
		default:
			err = guErrors.NewAppErrorWithType(guErrors.UnknownError)
		}
	}

	return
}
