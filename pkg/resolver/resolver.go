package resolver

import (
  "github.com/gomodule/redigo/redis"
  "github.com/kieranroneill/mimir/pkg/config"
  "gorm.io/gorm"
)

type Resolver struct {
  Config *config.Config
  Database *gorm.DB
  Store *redis.Pool
}
