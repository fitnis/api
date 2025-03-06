package handler

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "github.com/fitnis/api/internal/models"
)

type PatientHandler struct {
    DB *gorm.DB
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
    return &PatientHandler{DB: db}
}

// GET /patients
func (h *PatientHandler) ListPatients(c echo.Context) error {
    var patients []models.Patient

    // (Uncomment if you want to run actual queries)
    // err := h.DB.Find(&patients).Error
    // if err != nil {
    //     return c.JSON(http.StatusInternalServerError, map[string]string{
    //         "error": err.Error(),
    //     })
    // }

    // For now, return an empty array or a sample
    return c.JSON(http.StatusOK, patients)
}

// POST /patients
func (h *PatientHandler) CreatePatient(c echo.Context) error {
    var p models.Patient
    if err := c.Bind(&p); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid payload",
        })
    }

    // err := h.DB.Create(&p).Error
    // if err != nil {
    //     return c.JSON(http.StatusInternalServerError, map[string]string{
    //         "error": err.Error(),
    //     })
    // }

    return c.JSON(http.StatusCreated, p)
}
