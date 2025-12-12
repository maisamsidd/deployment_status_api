package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maisam9060/deployment_status/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	// Swagger UI route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	router.GET("/deployments", controllers.GetDeployments)
	router.POST("/addDeployment", controllers.PostDeployment)
}
