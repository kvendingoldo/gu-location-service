package locations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-location-service/controllers"
	"github.com/kvendingoldo/gu-location-service/model"
	"strconv"

	guErrors "github.com/kvendingoldo/gu-location-service/internal/errors"
	"net/http"
)

// GetLocation godoc
// @Tags location
// @Summary Search in some location within the provided radius.
// @Description Search for users in some location within the provided radius (with pagination).
// @Success 200 {object} int
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /location [get]
func SearchByRadius(c *gin.Context) {
	//:coordinates:radius

	//var users []model.User
	//
	//err := model.GetAllUsers(&users)
	//if err != nil {
	//	//appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
	//	//_ = c.Error(appError)
	//	return
	//}

	c.JSON(http.StatusOK, "")
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

	distance, err := model.GetDistance(uid, timeRange)
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
	if err := controllers.BindJSON(c, &req); err != nil {
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

	location := &model.Location{
		UID:         uid,
		Coordinates: req.Coordinates,
	}

	err = model.UpdateLocation(location)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, "")
}
