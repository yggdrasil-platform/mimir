package store

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

func New() *redis.Pool {
	port := "6379"
	if value, exists := os.LookupEnv("STORE_PORT"); exists {
		port = value
	}

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf(
					"%s:%s",
					os.Getenv("STORE_HOST"),
					port,
				),
				redis.DialPassword(os.Getenv("STORE_PASSWORD")),
			)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}
}
