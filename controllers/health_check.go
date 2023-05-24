package controllers

import (
	"database/sql"
	"net/http"
	"simple-gin-server/models"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct {
	Model *models.HealthCheckModel
}

func (healthCheckController HealthCheckController) ProbeLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}

func (healthCheckController HealthCheckController) ProbeReadiness(c *gin.Context) {
	err := healthCheckController.Model.CheckDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}

func NewHealthCheckController(database *sql.DB) *HealthCheckController {
	model := models.NewHealthCheckModel(database)
	return &HealthCheckController{Model: model}
}
