package handlers

import (
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(rg *gin.RouterGroup) {
	grp := rg.Group("/orders")
	grp.POST("/new", createOrder)
	grp.GET("/get", getOrders)
	grp.DELETE("/delete/:orderId", deleteOrder)
}

func createOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	c.JSON(http.StatusCreated, services.CreateOrder(req))
}

func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetOrders())
}

func deleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	services.DeleteOrder(id)
	c.Status(http.StatusNoContent)
}
