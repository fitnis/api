package handlers

import (
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterPrescriptionRoutes(rg *gin.RouterGroup) {
	p := rg.Group("/prescriptions")
	p.POST("", prescribe)
	p.POST("/prescribe", prescribe)
}

func prescribe(c *gin.Context) {
	var req models.Prescription
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.PrescribeMedication(req))
}
