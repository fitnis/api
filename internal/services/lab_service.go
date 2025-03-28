package services

import "github.com/fitnis/api/internal/models"

var samples []models.Sample
var evaluations []models.SampleEvaluation

func CollectSample(s models.Sample) models.GenericResponse {
	samples = append(samples, s)
	return models.GenericResponse{Message: "Sample collected"}
}

func RecordSample(s models.Sample) models.GenericResponse {
	return models.GenericResponse{Message: "Sample recorded"}
}

func ProcessSample(s models.Sample) models.GenericResponse {
	return models.GenericResponse{Message: "Sample processing started"}
}

func EvaluateSample(e models.SampleEvaluation) models.GenericResponse {
	evaluations = append(evaluations, e)
	return models.GenericResponse{Message: "Evaluation completed"}
}
