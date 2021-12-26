package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateUsername(username string) error {
	// TODO: check it
	//username - 4-16 symbols (a-zA-Z0-9 symbols are acceptable)
	return nil

}

func ValidateCoordinates(raw string) error {
	coordinates := strings.Split(raw, ",")
	if len(coordinates) == 2 {
		latitude, err := strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return errors.New("Failed to parse latitude")
		}
		if !(latitude < 90.0 && latitude > -90.0) {
			return errors.New("Wrong latitude coordinate; It should be between -90 and 90")
		}

		longitude, err := strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return errors.New("Failed to parse longitude")
		}

		if !(longitude < 180.0 && longitude > -180.0) {
			return errors.New("Wrong latitude coordinate; It should be between -180 and 180")
		}
	}

	return nil
}
