package config

import (
  "os"
)

type Config struct {
  ClientJWTSecretKey string
  EncryptionKey string
  Environment string
	Port string
  ServiceName string
  UserJWTSecretKey string
  Version string
}

func New() *Config {
	return &Config{
    ClientJWTSecretKey: GetEnv("CLIENT_JWT_SECRET_KEY", ""),
    EncryptionKey: GetEnv("ENCRYPTION_KEY", ""),
		Environment: GetEnv("ENV", ""),
		Port: GetEnv("PORT", ""),
    ServiceName: GetEnv("SERVICE_NAME", ""),
    UserJWTSecretKey: GetEnv("USER_JWT_SECRET_KEY", ""),
    Version: GetEnv("VERSION", ""),
	}
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
