package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func PrescribeHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","medication":"Ibuprofen","dosage":"200mg"}' http://localhost:8080/api/prescriptions/prescribe
	if r.Method == http.MethodPost {
		var req models.Prescription
		json.NewDecoder(r.Body).Decode(&req)
		services.PrescribeMedication(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func PrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","medication":"Paracetamol","dosage":"500mg"}' http://localhost:8080/api/prescriptions
	if r.Method == http.MethodPost {
		var req models.Prescription
		json.NewDecoder(r.Body).Decode(&req)
		services.CreatePrescription(req)
		w.WriteHeader(http.StatusCreated)
	}
}
