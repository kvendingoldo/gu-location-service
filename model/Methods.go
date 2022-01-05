package model

import (
	"encoding/json"
	"fmt"
	"github.com/kvendingoldo/gu-location-service/config"
	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"
	"github.com/kvendingoldo/gu-location-service/pkg/distance"
	"time"
)

func (location *Location) TableName() string {
	return "locations"
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

// SearchLocationsWithinRadius ... Search all users inside provided radius
func SearchLocationsWithinRadius(centerCoordinates string, radius float64, pagination *Pagination) ([]Location, error) {
	locations, err := GetLatestUserLocations(pagination)
	if err != nil {
		return []Location{}, err
	}

	result := []Location{}

	for _, location := range locations {
		coordinates := []string{
			centerCoordinates,
			location.Coordinates,
		}

		distanceKm, err := distance.CalculateDistanceInKm(coordinates)
		if err != nil {
			return []Location{}, err
		}

		if distanceKm <= radius {
			result = append(result, location)
		}
	}

	return result, nil
}

// GetLatestLocations ... Search all users inside of provided radius
func GetLatestUserLocations(pagination *Pagination) ([]Location, error) {
	var locations []Location

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	subQuery := config.Config.DB.
		Model(&Location{}).
		Select("uid, max(created_at) as created_at").Group("uid")

	err := queryBuider.Debug().
		Model(&Location{}).
		Select("last_data.uid, coordinates, last_data.created_at").
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
