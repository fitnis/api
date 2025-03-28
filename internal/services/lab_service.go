package services

import "github.com/fitnis/api/internal/models"

var samples []models.Sample
var evaluations []models.SampleEvaluation

func CollectSample(s models.Sample) {
	samples = append(samples, s)
}

func RecordSample(s models.Sample) {
	samples = append(samples, s)
}

func ProcessSample(s models.Sample) {
	samples = append(samples, s)
}

func EvaluateSample(e models.SampleEvaluation) {
	evaluations = append(evaluations, e)
}
