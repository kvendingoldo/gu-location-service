package distance

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coordinate struct {
	Latitude  float64
	Longitude float64
}

// CalculateDistanceInMiles calculates the distance between coordinates in input array in Miles
func CalculateDistanceInMiles(coordinates []string) (float64, error) {
	distance, err := calculateDistance(coordinates)
	if err != nil {
		return 0.0, err
	}
	return distance * 0.621371, nil
}

// CalculateDistanceInKm calculates the distance between coordinates in input array in Kilometers
func CalculateDistanceInKm(coordinates []string) (float64, error) {
	distance, err := calculateDistance(coordinates)
	if err != nil {
		return 0.0, err
	}
	return distance, nil
}

// calculateDistance calculates the distance between coordinates in input array in Kilometers
func calculateDistance(coordinatesRaw []string) (float64, error) {
	coordinates, err := getCoordinatesList(coordinatesRaw)
	if err != nil {
		return 0.0, err
	}

	distance := 0.0

	// start from 1 is not a mistake
	for i := 1; i < len(coordinates); i++ {
		distance += calculateDistanceBetweenTwoPoints(coordinates[i-1], coordinates[i])
	}

	return distance, nil
}

// calculateDistanceBetweenTwoPoints calculates distance between two coordinates in Kilometers
// https://en.wikipedia.org/wiki/Haversine_formula
// https://stackoverflow.com/questions/27928/calculate-distance-between-two-latitude-longitude-points-haversine-formula
func calculateDistanceBetweenTwoPoints(crd1, crd2 coordinate) float64 {
	diameterOfTheEarth := 12742.0
	p := math.Pi / 180
	a := 0.5 - math.Cos((crd2.Latitude-crd1.Latitude)*p)/2 + math.Cos(crd1.Latitude*p)*math.Cos(crd2.Latitude*p)*(1-math.Cos((crd2.Longitude-crd1.Longitude)*p))/2
	return diameterOfTheEarth * math.Asin(math.Sqrt(a))
}

// prepareCoordinatesList convert []string array of coordinates into []coordinate array
func getCoordinatesList(coordinatesRaw []string) ([]coordinate, error) {
	var coordinates []coordinate

	for _, coordinateRaw := range coordinatesRaw {
		coordinatePair := strings.Split(strings.ReplaceAll(coordinateRaw, " ", ""), ",")

		if len(coordinatePair) != 2 {
			return []coordinate{}, errors.New(fmt.Sprintf("Coordinate '%v' can not be parsed. It should be in format: 'c1, c2'", coordinatePair))
		}

		var latitude, longitude float64

		latitude, err := strconv.ParseFloat(coordinatePair[0], 64)
		if err != nil {
			return []coordinate{}, errors.New(fmt.Sprintf("Failed to parse latitude %v to float64", coordinatePair[0]))
		}

		longitude, err = strconv.ParseFloat(coordinatePair[1], 64)
		if err != nil {
			return []coordinate{}, errors.New(fmt.Sprintf("Failed to parse longitude %v to float64", coordinatePair[1]))
		}

		coordinates = append(coordinates, coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return coordinates, nil
}
