package handlers

import (
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterLabRoutes(rg *gin.RouterGroup) {
	lab := rg.Group("/lab")
	lab.POST("/collectSample", collect)
	lab.POST("/recordSample", record)
	lab.POST("/processSample", process)
	lab.POST("/evaluateSample", evaluate)
}

func collect(c *gin.Context) {
	var req models.Sample
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.CollectSample(req))
}

func record(c *gin.Context) {
	var req models.Sample
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.RecordSample(req))
}

func process(c *gin.Context) {
	var req models.Sample
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.ProcessSample(req))
}

func evaluate(c *gin.Context) {
	var req models.SampleEvaluation
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.EvaluateSample(req))
}
