package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// curl -X POST -d '{"name":"John Doe","dob":"1990-01-01","reason":"Chest pain"}' http://localhost:8080/api/patients/admit
		var req models.PatientRequest
		json.NewDecoder(r.Body).Decode(&req)
		services.AdmitPatient(req)
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		// curl http://localhost:8080/api/patients/admit
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services.ListAdmittedPatients())
	case http.MethodDelete:
		// curl -X DELETE "http://localhost:8080/api/patients/admit?patientId=123"
		id := r.URL.Query().Get("patientId")
		services.DischargePatient(id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"name":"John Doe","dob":"1990-01-01","reason":"Checkup"}' http://localhost:8080/api/patients/register
	if r.Method == http.MethodPost {
		var req models.PatientRequest
		json.NewDecoder(r.Body).Decode(&req)
		services.RegisterPatient(req)
		w.WriteHeader(http.StatusCreated)
	}
}
