package users

import (
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-user-service/controllers"
	"log"

	"github.com/kvendingoldo/gu-user-service/models"
	"net/http"
	"strconv"
)

// GetAllUsers godoc
// @Tags user
// @Summary Get all users
// @Description Get all users on the system
// @Success 200 {object} []models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User

	err := models.GetAllUsers(&users)
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
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{id} [get]
func GetUsersByID(c *gin.Context) {
	var user models.User

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	log.Println(userID)

	c.JSON(http.StatusOK, user)
}

// NewUser godoc
// @Tags user
// @Summary Create new user
// @Description Create new user on the system
// @Accept json
// @Produce json
// @Param data body NewUserRequest true "body data"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [post]
func NewUser(c *gin.Context) {
	var request NewUserRequest

	if err := controllers.BindJSON(c, &request); err != nil {
		//appError := errorModels.NewAppError(err, errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}
	user := models.User{
		Name: request.Name,
	}

	//err := models.CreateMedicine(&medicine)
	//if err != nil {
	//	_ = c.Error(err)
	//	return
	//}

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
	medicineID, _ := strconv.Atoi(c.Param("id"))
	log.Println(medicineID)
	//if err != nil {
	//	//appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
	//	//_ = c.Error(appError)
	//	return
	//}

	//err = models.DeleteMedicine(medicineID)
	//if err != nil {
	//	_ = c.Error(err)
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}