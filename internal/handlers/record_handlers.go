package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func ChartHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// curl -X POST -d '{"patientId":"123","note":"Patient is recovering well."}' http://localhost:8080/api/records/chart
		var req models.ChartNote
		json.NewDecoder(r.Body).Decode(&req)
		services.RecordChartNote(req)
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		// curl http://localhost:8080/api/records/chart
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services.GetChartNotes())
	}
}

func ExamHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","examType":"MRI"}' http://localhost:8080/api/records/exam
	if r.Method == http.MethodPost {
		var req models.ExamRequest
		json.NewDecoder(r.Body).Decode(&req)
		services.PerformExam(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func ExamResultsHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"patientId":"123","result":"No abnormalities found"}' http://localhost:8080/api/records/exam/results
	if r.Method == http.MethodPost {
		var req models.ExamResult
		json.NewDecoder(r.Body).Decode(&req)
		services.RecordExamResult(req)
		w.WriteHeader(http.StatusCreated)
	}
}
