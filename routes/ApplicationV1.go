package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kvendingoldo/gu-user-service/api"
	"github.com/kvendingoldo/gu-user-service/controllers/misc"
	"github.com/kvendingoldo/gu-user-service/controllers/users"
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

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// Misc
		v1Misc := v1.Group("/misc")
		{
			v1Misc.GET("/ping", misc.GetPing)
		}

		// Users
		v1Users := v1.Group("/users")
		{
			v1Users.GET("/", users.GetAllUsers)
			v1Users.GET("/:id", users.GetUsersByID)
			v1Users.POST("/", users.NewUser)
			v1Users.DELETE("/:id", users.DeleteUser)
		}
	}
}
