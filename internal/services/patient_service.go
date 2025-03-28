package services

import "github.com/fitnis/api/internal/models"

var patients []models.PatientRequest

func AdmitPatient(p models.PatientRequest) {
	patients = append(patients, p)
}

func ListAdmittedPatients() []models.PatientRequest {
	return patients
}

func DischargePatient(id string) {
	// mock discharge
}

func RegisterPatient(p models.PatientRequest) {
	// simulate registration
}
