package application

import (
  "github.com/gomodule/redigo/redis"
  "github.com/kieranroneill/mimir/pkg/config"
  "github.com/kieranroneill/mimir/pkg/database"
  "github.com/kieranroneill/mimir/pkg/logger"
  "github.com/kieranroneill/mimir/pkg/store"
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

  // Run db migrations
  if err = database.RunMigrations(db); err != nil {
    logger.Error.Printf("failed to run database migrations: %s", err)
  }

  // Create a Redis pool
  pool := store.New()

  return &Application{
    Config: config.New(),
    Database: db,
    Store: pool,
  }, nil
}
