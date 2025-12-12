package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maisam9060/deployment_status/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run("localhost:8090")

}
