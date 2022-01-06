package model

import (
	"encoding/json"
	"fmt"
	commonLibModels "github.com/kvendingoldo/gu-common/pkg/model"
	"github.com/kvendingoldo/gu-location-service/internal/distance"
	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"
	"time"
)

type Location struct {
	commonLibModels.Model
	UID int64 `json:"uid" example:"12"` // gorm:foreignKey

	Latitude  float64 `json:"lat" example:"39.12355"`
	Longitude float64 `json:"lon" example:"27.64538"`
}

// TableName represents name of SQL table, used by GORM
func (location *Location) TableName() string {
	return "locations"
}

// UpdateLocation ... Update location
func UpdateLocation(location *Location) (err error) {
	err = db.Create(location).Error
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

// SearchLocationsWithinRadius ... Search all users inside provided radius
func SearchLocationsWithinRadius(lat, lon, radiusInKm float64, pagination *commonLibModels.Pagination) ([]Location, error) {
	locations, err := GetLatestUserLocations(pagination)
	if err != nil {
		return []Location{}, err
	}

	result := []Location{}

	for _, location := range locations {
		coordinates := []distance.Coordinate{
			{Latitude: lat, Longitude: lon},
			{Latitude: location.Latitude, Longitude: location.Longitude},
		}

		distanceKm, err := distance.CalculateDistance(coordinates, "km")
		if err != nil {
			return []Location{}, err
		}

		if distanceKm <= radiusInKm {
			result = append(result, location)
		}
	}

	return result, nil
}

// GetLatestLocations ... Search all users inside of provided radius
func GetLatestUserLocations(pagination *commonLibModels.Pagination) ([]Location, error) {
	var locations []Location

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	subQuery := db.
		Model(&Location{}).
		Select("uid, max(created_at) as created_at").Group("uid")

	err := queryBuider.Debug().
		Model(&Location{}).
		Select("last_data.uid, latitude, longitude, last_data.created_at").
		Joins("inner join (?) as last_data on locations.uid = last_data.uid and locations.created_at = last_data.created_at", subQuery).
		Find(&locations).
		Error
	if err != nil {
		// TODO
		fmt.Println(err)
	}

	return locations, nil
}

// GetDistance ... Get user distance in Km by range
func GetDistance(uid int, timeRange string) (float64, error) {
	locations, err := getLocationsByRange(uid, timeRange)
	if err != nil {
		fmt.Println(err)
	}

	coordinates := []distance.Coordinate{}
	for _, location := range locations {
		coordinates = append(
			coordinates,
			distance.Coordinate{Latitude: location.Latitude, Longitude: location.Longitude},
		)
	}

	distanceKm, err := distance.CalculateDistance(coordinates, "km")
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

	db.Where(
		"created_at >= ? and uid = ?", time.Now().Add(delta), uid,
	).Find(&locations)

	return locations, nil
}
