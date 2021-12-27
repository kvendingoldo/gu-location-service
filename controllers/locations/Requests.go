package locations

import (
	"errors"
	"fmt"
	"github.com/kvendingoldo/gu-location-service/internal/utils"
)

type NewLocationRequest struct {
	UID         int64  `json:"uid,omitempty" example:"800"`
	Username    string `json:"username,omitempty" example:"Bill"`
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538" binding:"required"`
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

	err = utils.ValidateCoordinates(r.Coordinates)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

type NewSearchRequest struct {
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"  binding:"required"`
	Radius      string `json:"radius" example:"800" default:"100" binding:"optional,min=0"`
}

func (r *NewSearchRequest) validate() error {
	return utils.ValidateCoordinates(r.Coordinates)
}

type NewDistanceRequest struct {
	ID       string `json:"id,omitempty" example:"800"`
	Username string `json:"username,omitempty" example:"Bill"`
	Range    int    `json:"range" example:"1"  binding:"optional,min=0"`
}

func (r *NewDistanceRequest) validate() error {
	if r.ID == "" && r.Username == "" {
		return errors.New("At least one of id || username should be not empty")
	}

	err := utils.ValidateUsername(r.Username)
	if err != nil {
		return err
	}

	return nil
}