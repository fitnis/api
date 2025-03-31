package main

import (
	"github.com/fitnis/api/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Enable CORS

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		handlers.RegisterAppointmentRoutes(api)
		handlers.RegisterOrderRoutes(api)
		handlers.RegisterPatientRoutes(api)
		handlers.RegisterRecordRoutes(api)
		handlers.RegisterLabRoutes(api)
		handlers.RegisterPrescriptionRoutes(api)
		handlers.RegisterReferralRoutes(api)
	}

	router.Run(":8080")
}
