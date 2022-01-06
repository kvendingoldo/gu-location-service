package v1

import (
	"errors"
	"fmt"
	"github.com/kvendingoldo/gu-location-service/internal/utils"
)

type NewLocationRequest struct {
	UID      int64  `json:"uid,omitempty" example:"800"`
	Username string `json:"username,omitempty" example:"Bill"`

	Latitude  float64 `json:"lat" example:"39.12355" binding:"required"`
	Longitude float64 `json:"lon" example:"27.64538" binding:"required"`
}

func (r *NewLocationRequest) validate() error {
	if r.UID == 0 && r.Username == "" {
		return errors.New("At least one of id || username should be not empty")
	}

	fmt.Println("===")
	err := utils.ValidateUsername(r.Username)

	if err != nil {
		return err
	}

	err = utils.ValidateCoordinates(r.Latitude, r.Longitude)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
