package handlers

import (
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRecordRoutes(rg *gin.RouterGroup) {
	rec := rg.Group("/records")
	rec.POST("/chart", recordChart)
	rec.GET("/chart", getChart)
	rec.POST("/exam", performExam)
	rec.POST("/exam/results", recordExamResult)
}

func recordChart(c *gin.Context) {
	var req models.ChartNote
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.RecordChartNote(req))
}

func getChart(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetChartNotes())
}

func performExam(c *gin.Context) {
	var req models.ExamRequest
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.PerformExam(req))
}

func recordExamResult(c *gin.Context) {
	var req models.ExamResult
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.RecordExamResult(req))
}
