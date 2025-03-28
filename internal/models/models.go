package models

type AppointmentRequest struct {
	PatientID string `json:"patientId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Doctor    string `json:"doctor"`
}

type OrderRequest struct {
	PatientID string `json:"patientId"`
	Item      string `json:"item"`
	Priority  string `json:"priority"`
}

type PatientRequest struct {
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Reason string `json:"reason"`
}

type ChartNote struct {
	PatientID string `json:"patientId"`
	Note      string `json:"note"`
}

type ExamRequest struct {
	PatientID string `json:"patientId"`
	ExamType  string `json:"examType"`
}

type ExamResult struct {
	PatientID string `json:"patientId"`
	Result    string `json:"result"`
}

type Sample struct {
	SampleID  string `json:"sampleId"`
	PatientID string `json:"patientId"`
	Type      string `json:"type"`
}

type SampleEvaluation struct {
	SampleID string `json:"sampleId"`
	Result   string `json:"result"`
}

type Prescription struct {
	PatientID  string `json:"patientId"`
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
}

type Referral struct {
	PatientID  string `json:"patientId"`
	Department string `json:"department"`
	Reason     string `json:"reason"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
