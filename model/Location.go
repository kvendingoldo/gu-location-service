package model

import (
	"encoding/json"
	"fmt"
	"github.com/kvendingoldo/gu-location-service/config"
	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"
	"github.com/kvendingoldo/gu-location-service/pkg/distance"

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

// GetDistance ... Get user distance in Km by range
func GetDistance(uid int, timeRange string) (float64, error) {
	locations, err := getLocationsByRange(uid, timeRange)
	if err != nil {
		fmt.Println(err)
	}

	coordinates := []string{}
	for _, location := range locations {
		coordinates = append(coordinates, location.Coordinates)
	}

	distanceKm, err := distance.CalculateDistanceInKm(coordinates)
	if err != nil {
		fmt.Println(err)
	}

	return distanceKm, nil
}

// getLocationsByRange ... Get list of user's location by range
func getLocationsByRange(uid int, timeRange string) ([]Location, error) {
	var locations []Location

	delta, err := time.ParseDuration(timeRange)
	if err != nil {
		fmt.Println(err)
	}

	config.Config.DB.Where(
		"created_at >= ? and uid = ?", time.Now().Add(delta), uid,
	).Find(&locations)

	return locations, nil
}
