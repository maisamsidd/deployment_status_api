package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maisam9060/deployment_status/models"
)

func GetDeployments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Deployments)

}

func PostDeployment(c *gin.Context) {
	var newDeployment models.Deployment
	err := c.BindJSON(&newDeployment)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid JSON"})
		return
	}

	models.Deployments = append(models.Deployments, newDeployment)
	c.IndentedJSON(http.StatusCreated, newDeployment)
}
