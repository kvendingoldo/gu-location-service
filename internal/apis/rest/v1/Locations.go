package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"

	http2 "github.com/kvendingoldo/gu-location-service/internal/apis/rest"
	"github.com/kvendingoldo/gu-location-service/internal/models"
	"strconv"

	commonLibUtils "github.com/kvendingoldo/gu-common/pkg/utils"

	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"
	"net/http"
)

// GetLocation godoc
// @Tags location
// @Summary Search users in some location within the provided radius.
// @Description Search for users in some location within the provided radius (with pagination).
// @Param lat query string true "Center latitude"
// @Param lon query string true "Center longitude"
// @Param radius query number false "radius"
// @Param units query string false "radius (m|km|mi|ft)"
// @Success 200 {object} int
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [get]
func SearchByRadius(c *gin.Context) {
	// return only id!!!

	// Reference
	// http://localhost:8080/v1/users?lat=55.751508&lon=37.615666&radius=1&units=km
	// &page=1&limit=2

	reqLat := c.Query("lat")
	lat, err := strconv.ParseFloat(reqLat, 64)
	if err != nil {
		// TODO: error
	}

	reqLon := c.Query("lon")
	lon, err := strconv.ParseFloat(reqLon, 64)
	if err != nil {
		// TODO: error
	}

	radius := 0.1
	if reqRadius, ok := c.GetQuery("radius"); ok {
		units := "km"
		if reqUnits, ok := c.GetQuery("units"); ok {
			units = reqUnits
		}

		parsedRadius, err := strconv.ParseFloat(reqRadius, 64)
		if err != nil {
			// TODO: error
		}

		switch units {
		case "m":
			radius = parsedRadius * 0.1
		case "km":
			radius = parsedRadius
		case "ft":
			radius = parsedRadius * 3280.84
		case "mi":
			radius = parsedRadius * 0.621371
		default:
			// TODO: error
		}
	}

	pagination := commonLibUtils.GeneratePaginationFromRequest(c)
	locations, err := models.SearchLocationsWithinRadius(lat, lon, radius, &pagination)
	if err != nil {
		// TODO: error
	}

	uids := []int64{}
	for _, location := range locations {
		uids = append(uids, location.UID)
	}

	c.JSON(http.StatusOK, uids)
}

// TODO
// name type datatype mandatory comment

// GetDistance godoc
// @Tags location
// @Summary Returns distance traveled by a person within some date/time range (in days).
// @Description Returns distance traveled by a person within some date/time range (in days).
// @Param uid query int true "id of user"
// @Param range query string false "time range"
// @Success 200 {object} int
// TODO @Failure 400 {object} MessageResponse
// TODO @Failure 500 {object} MessageResponse
// @Router /distance [get]
func GetDistance(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		fmt.Println(err)
		return
	}

	timeRange := "-1h"
	if reqTimeRange, ok := c.GetQuery("range"); ok {
		timeRange = reqTimeRange
	}

	distance, err := models.GetDistance(uid, timeRange)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, distance)
}

// UpdateLocation godoc
// @Tags location
// @Summary Update current user location by the username/uid.
// @Description Update current user location by the username/uid.
// @Param data body NewLocationRequest true "body data"
// @Success 200 {string} string	"ok"
// @Router /location [put]
func UpdateLocation(c *gin.Context) {
	var req NewLocationRequest
	if err := http2.BindJSON(c, &req); err != nil {
		appError := guErrors.NewAppError(err, guErrors.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := req.validate()
	if err != nil {
		appError := guErrors.NewAppError(err, guErrors.ValidationError)
		_ = c.Error(appError)
		return
	}

	var uid int64
	if req.UID == 0 {
		// todo
		uid = 0
	} else {
		uid = req.UID
	}

	location := &models.Location{
		UID:       uid,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	err = models.UpdateLocation(location)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, "")
}
