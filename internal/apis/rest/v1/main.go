package v1

import (
	"github.com/gin-gonic/gin"
	appErrors "github.com/kvendingoldo/gu-common/pkg/errors"
	appUtils "github.com/kvendingoldo/gu-common/pkg/utils"
	"github.com/kvendingoldo/gu-user-service/internal/models"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
)

// @title GU user service
// @version 1.0
// @description Documentation's GU user service
// @termsOfService http://swagger.io/terms/
// @contact.name Alexander Sharov
// @contact.url http://github.com/kvendingoldo
// @contact.email kvendingoldo@gmail.com
// @BasePath /v1
// @accept json
// @produce json
// @schemes http

func NewRouter(router *gin.Engine) {

	//api.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%v", config.Config.RestPort)

	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// Users
		v1Users := v1.Group("/users")
		{
			v1Users.POST("/", NewUser)
			v1Users.GET("/", GetAllUsers)
			v1Users.GET("/:id", GetUserByID)
			v1Users.PUT("/:id", UpdateUser)
			v1Users.DELETE("/:id", DeleteUser)
		}
	}
}

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
func GetUserByID(c *gin.Context) {
	var user models.User

	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	err = models.GetUserByID(&user, userID)
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
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [post]
func NewUser(c *gin.Context) {
	var req NewUserRequest

	if err := appUtils.BindJSON(c, &req); err != nil {
		appError := appErrors.NewAppError(err, appErrors.ValidationError)
		_ = c.Error(appError)
		return
	}

	user := models.User{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	err := models.CreateUser(&user)
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
// @Param input body models.User true "User updated info"
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

	err = appUtils.BindJSONMap(c, &requestMap)
	if err != nil {
		//appError := errorModels.NewAppError(err, errorModels.ValidationError)
		//_ = c.Error(appError)
		return
	}

	user, err := models.UpdateUser(userID, requestMap)
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

	err = models.DeleteUser(userID)
	if err != nil {
		//_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
