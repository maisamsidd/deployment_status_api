package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maisam9060/deployment_status/controllers"
)

func SetupRoutes(routes *gin.Engine) {
	routes.GET("/deployments", controllers.GetDeployments)
	routes.POST("/addDeployment", controllers.PostDeployment)
}
