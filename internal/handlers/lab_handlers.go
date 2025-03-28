package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fitnis/api/internal/models"
	"github.com/fitnis/api/internal/services"
)

func CollectSampleHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"sampleId":"s1","patientId":"123","type":"blood"}' http://localhost:8080/api/lab/collectSample
	if r.Method == http.MethodPost {
		var req models.Sample
		json.NewDecoder(r.Body).Decode(&req)
		services.CollectSample(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func RecordSampleHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"sampleId":"s2","patientId":"123","type":"urine"}' http://localhost:8080/api/lab/recordSample
	if r.Method == http.MethodPost {
		var req models.Sample
		json.NewDecoder(r.Body).Decode(&req)
		services.RecordSample(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func ProcessSampleHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"sampleId":"s3","patientId":"123","type":"blood"}' http://localhost:8080/api/lab/processSample
	if r.Method == http.MethodPost {
		var req models.Sample
		json.NewDecoder(r.Body).Decode(&req)
		services.ProcessSample(req)
		w.WriteHeader(http.StatusCreated)
	}
}

func EvaluateSampleHandler(w http.ResponseWriter, r *http.Request) {
	// curl -X POST -d '{"sampleId":"s4","result":"Negative"}' http://localhost:8080/api/lab/evaluateSample
	if r.Method == http.MethodPost {
		var req models.SampleEvaluation
		json.NewDecoder(r.Body).Decode(&req)
		services.EvaluateSample(req)
		w.WriteHeader(http.StatusCreated)
	}
}
