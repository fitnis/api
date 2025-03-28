package services

import "github.com/fitnis/api/internal/models"

var orders []models.OrderRequest

func CreateOrder(req models.OrderRequest) models.GenericResponse {
	orders = append(orders, req)
	return models.GenericResponse{Message: "Order created"}
}

func GetOrders() []models.OrderRequest {
	return orders
}

func DeleteOrder(id string) {
	orders = []models.OrderRequest{}
}
