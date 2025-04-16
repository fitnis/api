// filepath: cmd/api/main.go
package main

import (
	"log"

	"github.com/fitnis/api/internal/database" // Import database package
	"github.com/fitnis/api/internal/handlers"
	"github.com/fitnis/api/internal/services" // Import services package
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.InitDB()
	db := database.DB // Get the DB instance

	// Initialize Services with DB instance
	patientService := services.NewPatientService(db)
	examinationService := services.NewExaminationService(db)
	sampleService := services.NewSampleService(db, services.NewPrescriptionService(db)) // Sample service needs prescription service
	prescriptionService := services.NewPrescriptionService(db)
	referralService := services.NewReferralService(db)

	// Initialize Handlers with Services
	patientHandler := handlers.NewPatientHandler(patientService)
	examinationHandler := handlers.NewExaminationHandler(examinationService)
	sampleHandler := handlers.NewSampleHandler(sampleService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService)
	referralHandler := handlers.NewReferralHandler(referralService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		// Patient endpoints
		patients := api.Group("/patients")
		{
			patients.GET("", patientHandler.GetPatients)
			patients.GET("/:id", patientHandler.GetPatient)
			patients.POST("", patientHandler.CreatePatient)
			patients.PUT("/:id", patientHandler.UpdatePatient)
			patients.DELETE("/:id", patientHandler.DeletePatient)
		}

		// Examination endpoints
		examinations := api.Group("/examinations")
		{
			examinations.GET("", examinationHandler.GetExaminations)
			examinations.GET("/:id", examinationHandler.GetExamination)
			examinations.POST("", examinationHandler.CreateExamination)
			examinations.PUT("/:id", examinationHandler.UpdateExamination)
			examinations.DELETE("/:id", examinationHandler.DeleteExamination)
			examinations.GET("/patient/:patientId", examinationHandler.GetExaminationsByPatientID)
		}

		// Sample endpoints
		samples := api.Group("/samples")
		{
			samples.GET("", sampleHandler.GetSamples)
			samples.GET("/:id", sampleHandler.GetSample)
			samples.POST("", sampleHandler.CreateSample)
			samples.PUT("/:id", sampleHandler.UpdateSample)
			samples.DELETE("/:id", sampleHandler.DeleteSample)
			samples.GET("/examination/:examinationId", sampleHandler.GetSamplesByExaminationID)
		}

		// Prescription endpoints
		prescriptions := api.Group("/prescriptions")
		{
			prescriptions.GET("", prescriptionHandler.GetPrescriptions)
			prescriptions.GET("/:id", prescriptionHandler.GetPrescription)
			prescriptions.POST("", prescriptionHandler.CreatePrescription)
			prescriptions.PUT("/:id", prescriptionHandler.UpdatePrescription)
			prescriptions.DELETE("/:id", prescriptionHandler.DeletePrescription)
			prescriptions.GET("/examination/:examinationId", prescriptionHandler.GetPrescriptionsByExaminationID)
			prescriptions.POST("/:id/validate", prescriptionHandler.ValidatePrescription)
			prescriptions.POST("/:id/send", prescriptionHandler.SendPrescription)
		}

		// Referral endpoints
		referrals := api.Group("/referrals")
		{
			referrals.GET("", referralHandler.GetReferrals)
			referrals.GET("/:id", referralHandler.GetReferral)
			referrals.POST("", referralHandler.CreateReferral)
			referrals.PUT("/:id", referralHandler.UpdateReferral)
			referrals.DELETE("/:id", referralHandler.DeleteReferral)
			referrals.GET("/examination/:examinationId", referralHandler.GetReferralsByExaminationID)
		}
	}

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
