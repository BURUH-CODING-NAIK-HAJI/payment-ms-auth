package myredis

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	if client == nil {
		err := errors.New("FAILED CONNECT TO REDIS")
		panic(err)
	}

	return client
}
