package distance

import (
	"math"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// CalculateDistance calculates the distance between coordinates in input array in `units`
func CalculateDistance(coordinates []Coordinate, units string) (float64, error) {
	distance, err := calculateDistance(coordinates)
	if err != nil {
		return 0.0, err
	}

	switch units {
	case "m":
		distance = distance * 1000
	case "km":
		// TODO
	case "ft":
		distance = distance * 3280.84
	case "mi":
		distance = distance * 0.621371
	default:
		// TODO: error
	}

	return distance, nil
}

// calculateDistance calculates the distance between coordinates in input array in Kilometers
func calculateDistance(coordinates []Coordinate) (float64, error) {
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
func calculateDistanceBetweenTwoPoints(crd1, crd2 Coordinate) float64 {
	diameterOfTheEarth := 12742.0
	p := math.Pi / 180
	a := 0.5 - math.Cos((crd2.Latitude-crd1.Latitude)*p)/2 + math.Cos(crd1.Latitude*p)*math.Cos(crd2.Latitude*p)*(1-math.Cos((crd2.Longitude-crd1.Longitude)*p))/2
	return diameterOfTheEarth * math.Asin(math.Sqrt(a))
}
