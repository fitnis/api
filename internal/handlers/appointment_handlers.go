package handlers

import (
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(rg *gin.RouterGroup) {
	grp := rg.Group("/appointments")
	grp.POST("/schedule", scheduleAppointment)
	grp.GET("/schedule", getAppointments)
	grp.DELETE("/schedule/:appointmentId", cancelAppointment)
}

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","date":"2025-04-01","time":"14:00","doctor":"Dr. House"}' http://localhost:8080/api/appointments/schedule
func scheduleAppointment(c *gin.Context) {
	var req models.AppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	c.JSON(http.StatusCreated, services.ScheduleAppointment(req))
}

func getAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAppointments())
}

func cancelAppointment(c *gin.Context) {
	id := c.Param("appointmentId")
	services.CancelAppointment(id)
	c.Status(http.StatusNoContent)
}
