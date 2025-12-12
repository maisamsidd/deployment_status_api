package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maisam9060/deployment_status/models"
)

// @Summary Get all deployments
// @Description Get list of all deployments
// @Tags deployments
// @Produce json
// @Success 200 {array} models.Deployment
// @Router /deployments [get]
func GetDeployments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Deployments)

}

// @Summary Add new deployment
// @Description Create a new deployment
// @Tags deployments
// @Accept json
// @Produce json
// @Param deployment body models.Deployment true "Deployment object"
// @Success 201 {object} models.Deployment
// @Failure 400 {object} map[string]string
// @Router /addDeployment [post]
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
