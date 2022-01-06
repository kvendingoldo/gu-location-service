package utils

import (
	"errors"
)

func ValidateUsername(username string) error {
	// TODO: check it
	//username - 4-16 symbols (a-zA-Z0-9 symbols are acceptable)
	return nil

}

func ValidateCoordinates(latitude, longitude float64) error {
	if !(latitude < 90.0 && latitude > -90.0) {
		return errors.New("wrong latitude coordinate; It should be between -90 and 90")
	}

	if !(longitude < 180.0 && longitude > -180.0) {
		return errors.New("wrong latitude coordinate; It should be between -180 and 180")
	}

	return nil
}
