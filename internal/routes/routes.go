package routes

import (
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"

    "github.com/fitnis/api/internal/handler"
    "github.com/fitnis/api/internal/middleware"
)

func Register(e *echo.Echo, db *gorm.DB) {
    // Group routes under /api
    api := e.Group("/api")

    // Require authentication for all /api/* endpoints
    api.Use(middleware.DummyAuthMiddleware)

    // Create a patient handler
    ph := handler.NewPatientHandler(db)

    // Example routes
    api.GET("/patients", ph.ListPatients)
    api.POST("/patients", ph.CreatePatient)
}
