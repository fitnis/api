package services

import "github.com/fitnis/api/internal/models"

var notes []models.ChartNote
var exams []models.ExamRequest
var results []models.ExamResult

func RecordChartNote(note models.ChartNote) {
	notes = append(notes, note)
}

func GetChartNotes() []models.ChartNote {
	return notes
}

func PerformExam(exam models.ExamRequest) {
	exams = append(exams, exam)
}

func RecordExamResult(res models.ExamResult) {
	results = append(results, res)
}
