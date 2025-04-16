package models

import (
	"time"
)

// Patient model
type Patient struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	BirthDate time.Time
	Details   string

	// One-to-many relationship: a patient can have multiple examinations
	Examinations []Examination
}

// Examination model
type Examination struct {
	ID        uint      `gorm:"primaryKey"`
	PatientID uint      // foreign key for Patient
	ExamDate  time.Time `gorm:"not null"`
	Anamnesis string
	Diagnosis string

	// Belongs to
	Patient Patient

	// One-to-many relationships
	Samples       []Sample
	Prescriptions []Prescription
	Referrals     []Referral
}

// Sample model
type Sample struct {
	ID            uint `gorm:"primaryKey"`
	ExaminationID uint // foreign key for Examination
	SampleType    string
	Result        string

	// Belongs to
	Examination Examination
}

// Prescription model
type Prescription struct {
	ID            uint `gorm:"primaryKey"`
	ExaminationID uint // foreign key for Examination
	Medication    string
	Dosage        string
	Instructions  string
	Validated     bool
	Sent          bool

	// Belongs to
	Examination Examination
}

// Referral model
type Referral struct {
	ID            uint `gorm:"primaryKey"`
	ExaminationID uint // foreign key for Examination
	Specialist    string
	Reason        string

	// Belongs to
	Examination Examination
}
