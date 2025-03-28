package main

import (
	"log"
	"net/http"

	"github.com/fitnis/api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Scheduling
	mux.HandleFunc("/api/appointments/schedule", handlers.AppointmentHandler)
	mux.HandleFunc("/api/orders", handlers.OrderHandler)

	// Patients
	mux.HandleFunc("/api/patients/admit", handlers.PatientHandler)
	mux.HandleFunc("/api/patients/register", handlers.RegisterHandler)

	// Records
	mux.HandleFunc("/api/records/chart", handlers.ChartHandler)
	mux.HandleFunc("/api/records/exam", handlers.ExamHandler)
	mux.HandleFunc("/api/records/exam/results", handlers.ExamResultsHandler)

	// Lab
	mux.HandleFunc("/api/lab/collectSample", handlers.CollectSampleHandler)
	mux.HandleFunc("/api/lab/recordSample", handlers.RecordSampleHandler)
	mux.HandleFunc("/api/lab/processSample", handlers.ProcessSampleHandler)
	mux.HandleFunc("/api/lab/evaluateSample", handlers.EvaluateSampleHandler)

	// Prescriptions
	mux.HandleFunc("/api/prescriptions", handlers.PrescriptionHandler)
	mux.HandleFunc("/api/prescriptions/prescribe", handlers.PrescribeHandler)

	// Referrals
	mux.HandleFunc("/api/referrals", handlers.ReferralHandler)
	mux.HandleFunc("/api/referrals/specialist", handlers.SpecialistReferralHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
