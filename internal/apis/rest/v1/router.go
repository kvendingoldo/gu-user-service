package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kvendingoldo/gu-user-service/swagger_gen/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GU user service
// @version 1.0
// @description Documentation's GU user service
// @termsOfService http://swagger.io/terms/

// @contact.name Alexander Sharov
// @contact.url http://github.com/kvendingoldo
// @contact.email kvendingoldo@gmail.com

// @host localhost:8080
// @BasePath /v1
func NewRouter(router *gin.Engine) {

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
