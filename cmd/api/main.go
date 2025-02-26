package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize a new Echo instance
	e := echo.New()

	// Define a simple GET endpoint at "/"
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello from Echo!",
		})
	})

	// Start the server on port 3000
	log.Fatal(e.Start(":3000"))
}
