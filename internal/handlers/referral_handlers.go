package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func SpecialistReferralHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","department":"Cardiology","reason":"Chest pain"}' http://localhost:8080/api/referrals/specialist
	if r.Method == http.MethodPost {
		var req models.Referral
		json.NewDecoder(r.Body).Decode(&req)
		services.ReferToSpecialist(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func ReferralHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","department":"Neurology","reason":"Seizures"}' http://localhost:8080/api/referrals
	if r.Method == http.MethodPost {
		var req models.Referral
		json.NewDecoder(r.Body).Decode(&req)
		services.CreateReferral(req)
		w.WriteHeader(http.StatusCreated)
	}
}
