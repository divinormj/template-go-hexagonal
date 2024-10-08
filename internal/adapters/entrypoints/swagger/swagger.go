// Package swagger provides the setup for the Swagger documentation.
package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupSwagger sets up the Swagger documentation.
// @title			Template API with Swagger
// @version		1.0
// @description	This is an example API with Swagger using Gin.
// @host			localhost:8080
// @BasePath		/
func SetupSwagger(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
