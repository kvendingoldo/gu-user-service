package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-user-service/controllers"
	"github.com/kvendingoldo/gu-user-service/model"
	"github.com/kvendingoldo/gu-user-service/model/errors"
	"net/http"
	"strconv"
)

// GetAllUsers godoc
// @Tags user
// @Summary Get all users
// @Description Get all users on the system
// @Success 200 {object} []model.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []model.User

	err := model.GetAllUsers(&users)
	if err != nil {
		//appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
		//_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUsersByID godoc
// @Tags user
// @Summary Get users by ID
// @Description Get users by ID on the system
// @Param id path int true "id of user"
// @Success 200 {object} model.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	var user model.User

	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	err = model.GetUserByID(&user, userID)
	if err != nil {
		//appError := errorModels.NewAppError(err, errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// NewUser godoc
// @Tags user
// @Summary Create new user
// @Description Create new user on the system
// @Accept json
// @Produce json
// @Param data body NewUserRequest true "body data"
// @Success 200 {object} model.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [post]
func NewUser(c *gin.Context) {
	var req NewUserRequest

	if err := controllers.BindJSON(c, &req); err != nil {
		appError := errors.NewAppError(err, errors.ValidationError)
		_ = c.Error(appError)
		return
	}

	user := model.User{
		Name:        req.Name,
		Coordinates: req.Coordinates,
	}

	err := model.CreateUser(&user)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Tags user
// @Summary Update user
// @Description Update user on the system
// @Param id path int true "id of user"
// @Param input body model.User true "User updated info"
// @Success 200 {string} string	"ok"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		//appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}
	var requestMap map[string]interface{}

	err = controllers.BindJSONMap(c, &requestMap)
	if err != nil {
		//appError := errorModels.NewAppError(err, errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}

	err = updateValidation(requestMap)
	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err := model.UpdateUser(userID, requestMap)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Tags user
// @Summary Delete user
// @Description Delete user on the system
// @Param id path int true "id of user"
// @Success 200 {string} string	"ok"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		//appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}

	err = model.DeleteUser(userID)
	if err != nil {
		//_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
