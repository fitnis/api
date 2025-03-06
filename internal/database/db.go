package database

import (
	"fmt"

	"github.com/fitnis/api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	pg_migrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	driver, err := pg_migrate.WithInstance(sqlDB, &pg_migrate.Config{})
	m, m_err := migrate.NewWithDatabaseInstance("file://../layered-infra/migrations", "postgres", driver)

	if m_err != nil {
		panic(m_err)
	}

	m.Up()

	return db, nil
}
