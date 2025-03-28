package services

import "github.com/fitnis/api/internal/models"

var referrals []models.Referral

func ReferToSpecialist(r models.Referral) models.GenericResponse {
	referrals = append(referrals, r)
	return models.GenericResponse{Message: "Referred to specialist"}
}
