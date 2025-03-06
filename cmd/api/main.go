package main

import (
    "log"

    "github.com/fitnis/api/internal/config"
    "github.com/fitnis/api/internal/database"
    "github.com/fitnis/api/internal/routes"

    "github.com/labstack/echo/v4"
)

func main() {
    // 1. Load Viper config
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // 2. Initialize DB
    db, err := database.InitDB(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    // 3. Create Echo instance
    e := echo.New()

    // 4. Register routes
    routes.Register(e, db)

    // 5. Start server on configured port or default
    port := cfg.ServerPort
    if port == "" {
        port = "3000"
    }
    log.Fatal(e.Start(":" + port))
}
