package application

import (
  "github.com/gomodule/redigo/redis"
  "github.com/kieranroneill/new-go-service-template/pkg/config"
  "github.com/kieranroneill/new-go-service-template/pkg/database"
  "github.com/kieranroneill/new-go-service-template/pkg/store"
  "gorm.io/gorm"
)

type Application struct {
	Config *config.Config
  Database *gorm.DB
  Store *redis.Pool
}

func New() (*Application, error) {
  // Connect to the DB
  db, err := database.New()
  if err != nil {
    return nil, err
  }

  // Create a Redis pool
  pool := store.New()

  return &Application{
    Config: config.New(),
    Database: db,
    Store: pool,
  }, nil
}
