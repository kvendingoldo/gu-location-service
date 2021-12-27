package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-location-service/controllers/locations"
	_ "github.com/kvendingoldo/gu-location-service/swagger_gen/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GU location service
// @version 1.0
// @description Documentation's GU location service
// @termsOfService http://swagger.io/terms/

// @contact.name Alexander Sharov
// @contact.url http://github.com/kvendingoldo
// @contact.email kvendingoldo@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func ApplicationRouter(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		// Locations
		v1Locations := v1.Group("/")
		{
			v1Locations.GET("/distance", locations.GetDistance)
			v1Locations.GET("/search", locations.SearchByRadius)
			v1Locations.PUT("/location", locations.UpdateLocation)
		}
	}
}
