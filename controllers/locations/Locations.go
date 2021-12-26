package locations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-location-service/controllers"
	"github.com/kvendingoldo/gu-location-service/model/errors"
	"net/http"
	"strconv"
)

// GetLocation godoc
// @Tags location
// @Summary Get location
// @Description Get all users on the system
// @Success 200 {object} int
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [get]
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

// GetDistance godoc
// @Tags location
// @Summary Get user distance by range
// @Description TODO
// @Param data body NewDistanceRequest true "body data"
// @Success 200 {object} int
// TODO @Failure 400 {object} MessageResponse
// TODO @Failure 500 {object} MessageResponse
// @Router /distance [get]
func GetDistance(c *gin.Context) {
	//:id:range

	//var req NewDistanceRequest
	//
	//if err := controllers.BindJSON(c, &req); err != nil {
	//	//appError := errors.NewAppError(err, errors.ValidationError)
	//	//_ = c.Error(appError)
	//	return
	//}
	//
	//err := updateValidation(requestMap)
	//if err != nil {
	//	//_ = c.Error(err)
	//	return
	//}
	//
	////default range 1 day
	//var distance int
	//
	//userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	//if err != nil {
	//	return
	//}
	//
	//err = model.GetUserByID(&user, userID)
	//if err != nil {
	//	//appError := errorModels.NewAppError(err, errorModels.ValidationError)
	//	//_ = c.Error(appError)
	//	return
	//}

	c.JSON(http.StatusOK, "")
}

// UpdateLocation godoc
// @Tags location
// @Summary Update location
// @Description Update location on the system
// @Param data body NewLocationRequest true "body data"
// @Success 200 {string} string	"ok"
// @Router /location [put]
func UpdateLocation(c *gin.Context) {
	var req NewLocationRequest
	if err := controllers.BindJSON(c, &req); err != nil {
		appError := errors.NewAppError(err, errors.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := req.validate()
	if err != nil {
		// TODO
		fmt.Println("ERROR VALIDATE")

	}

	var uid int64
	//:id::username:coordinates

	uidRaw := c.Query("uid")
	if uidRaw != "" {
		var err error
		uid, err = strconv.ParseInt(uidRaw, 10, 64)
		if err != nil {
			//appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
			//_ = c.Error(appError)
			fmt.Println("eror")
			fmt.Println(err)
			return
		}
	}

	username := c.Query("username")
	coordinates := c.Query("coordinates")

	if uid == 0 {
		// make req to user service
		// get uid
	}

	fmt.Println(uid)
	fmt.Println(username)
	fmt.Println(coordinates)

	//err = model.UpdateLocation(uid, coordinates)
	//if err != nil {
	//	//_ = c.Error(err)
	//	return
	//}

	c.JSON(http.StatusOK, "")
}
