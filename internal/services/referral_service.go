package services

import "github.com/fitnis/api/internal/models"

var referrals []models.Referral

func ReferToSpecialist(r models.Referral) {
	referrals = append(referrals, r)
}

func CreateReferral(r models.Referral) {
	referrals = append(referrals, r)
}
