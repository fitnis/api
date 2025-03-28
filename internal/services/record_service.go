package services

import "github.com/fitnis/api/internal/models"

var chartNotes []models.ChartNote
var examRequests []models.ExamRequest
var examResults []models.ExamResult

func RecordChartNote(note models.ChartNote) models.GenericResponse {
	chartNotes = append(chartNotes, note)
	return models.GenericResponse{Message: "Chart updated"}
}

func GetChartNotes() []models.ChartNote {
	return chartNotes
}

func PerformExam(req models.ExamRequest) models.GenericResponse {
	examRequests = append(examRequests, req)
	return models.GenericResponse{Message: "Exam performed"}
}

func RecordExamResult(result models.ExamResult) models.GenericResponse {
	examResults = append(examResults, result)
	return models.GenericResponse{Message: "Result recorded"}
}
