package controllers

import (
	"errors"
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

	if err := c.BindJSON(&newDeployment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	// Insert at 0th index (latest first)
	models.Deployments = append(
		[]models.Deployment{newDeployment},
		models.Deployments...,
	)

	c.IndentedJSON(http.StatusCreated, newDeployment)
}

func getDeploymentId(id string) (*models.Deployment, error) {

	for i, d := range models.Deployments {
		if d.Id == id {
			return &models.Deployments[i], nil
		}
	}
	return nil, errors.New("invalid id")

}

func GetDeploymentsById(c *gin.Context) {
	id := c.Param("id")
	deployment, err := getDeploymentId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "deployment not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, deployment)
}
