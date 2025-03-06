package middleware

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func DummyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Unauthorized - missing token",
            })
        }
        return next(c)
    }
}
