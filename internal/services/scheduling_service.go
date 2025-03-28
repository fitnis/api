package services

import "github.com/fitnis/api/internal/models"

var appointments []models.AppointmentRequest
var orders []models.OrderRequest

func ScheduleAppointment(req models.AppointmentRequest) models.GenericResponse {
	appointments = append(appointments, req)
	return models.GenericResponse{Message: "Appointment scheduled"}
}

func GetAppointments() []models.AppointmentRequest {
	return appointments
}

func CancelAppointment(id string) models.GenericResponse {
	// mock cancel
	return models.GenericResponse{Message: "Appointment cancelled"}
}

func CreateOrder(req models.OrderRequest) {
	orders = append(orders, req)
}

func GetOrders() []models.OrderRequest {
	return orders
}

func DeleteOrder(id string) models.GenericResponse {
	// mock delete, no-op
	return models.GenericResponse{Message: "Order deleted"}
}
