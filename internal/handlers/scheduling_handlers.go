package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func AppointmentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// curl -X POST -d '{"patientId":"123","date":"2025-04-01","time":"14:00","doctor":"Dr. House"}' http://localhost:8080/api/appointments/schedule
		var req models.AppointmentRequest
		json.NewDecoder(r.Body).Decode(&req)
		resp := services.ScheduleAppointment(req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	case http.MethodGet:
		// curl http://localhost:8080/api/appointments/schedule
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services.GetAppointments())
	case http.MethodDelete:
		// curl -X DELETE "http://localhost:8080/api/appointments/schedule?appointmentId=abc"
		id := r.URL.Query().Get("appointmentId")
		services.CancelAppointment(id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// curl -X POST -d '{"patientId":"123","item":"X-ray","priority":"high"}' http://localhost:8080/api/orders
		var req models.OrderRequest
		json.NewDecoder(r.Body).Decode(&req)
		services.CreateOrder(req)
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		// curl http://localhost:8080/api/orders
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services.GetOrders())
	case http.MethodDelete:
		// curl -X DELETE "http://localhost:8080/api/orders?orderId=xyz"
		id := r.URL.Query().Get("orderId")
		services.DeleteOrder(id)
		w.WriteHeader(http.StatusNoContent)
	}
}
