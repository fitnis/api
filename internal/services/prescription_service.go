package services

import "github.com/fitnis/api/internal/models"

var prescriptions []models.Prescription

func PrescribeMedication(p models.Prescription) {
	prescriptions = append(prescriptions, p)
}

func CreatePrescription(p models.Prescription) {
	prescriptions = append(prescriptions, p)
}
