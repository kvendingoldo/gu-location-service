package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateUsername(username string) error {
	// TODO: check it
	//username - 4-16 symbols (a-zA-Z0-9 symbols are acceptable)
	return nil

}

func ValidateCoordinates(raw string) error {
	coordinates := strings.Split(strings.ReplaceAll(raw, " ", ""), ",")
	if len(coordinates) == 2 {
		latitude, err := strconv.ParseFloat(coordinates[0], 64)
		if err != nil {
			return errors.New("failed to parse latitude")
		}
		if !(latitude < 90.0 && latitude > -90.0) {
			return errors.New("wrong latitude coordinate; It should be between -90 and 90")
		}

		longitude, err := strconv.ParseFloat(coordinates[1], 64)
		if err != nil {
			fmt.Println(err)
			return errors.New("failed to parse longitude")
		}

		if !(longitude < 180.0 && longitude > -180.0) {
			return errors.New("wrong latitude coordinate; It should be between -180 and 180")
		}
	} else {
		return errors.New("wrong number of coordinates; Please pass coordinates in format: 'latitude,longitude'")
	}

	return nil
}
