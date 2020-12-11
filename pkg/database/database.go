package database

import (
	"fmt"
	"github.com/kieranroneill/new-go-service-template/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func New() (*gorm.DB, error) {
	dbPort := "5432"
	if value, exists := os.LookupEnv("DB_PORT"); exists {
		dbPort = value
	}

	dsn := fmt.Sprintf(
		"dbname=%s host=%s password=%s port=%s user=%s sslmode=disable",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PASSWORD"),
		dbPort,
		os.Getenv("DB_USER"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Print(err)
		return nil, err
	}

	return db, nil
}
