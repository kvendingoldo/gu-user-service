package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/swagger_gen/api"
	"strconv"

	"github.com/kvendingoldo/gu-user-service/internal/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title GUS user service
// @version 2.0
// @description Documentation's GU user service
// @termsOfService http://swagger.io/terms/
// @contact.name Alexander Sharov
// @contact.url http://github.com/kvendingoldo
// @contact.email kvendingoldo@gmail.com
// @BasePath /v2
// @accept json
// @produce json
// @schemes http

func NewRouter(router *gin.Engine) {

	api.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%v", config.Config.RestPort)

	v2 := router.Group("/v2")
	{
		// Documentation Swagger
		{
			v2.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// Users
		v2Users := v2.Group("/users")
		{
			v2Users.DELETE("/:id", DeleteUser)
		}
	}
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
