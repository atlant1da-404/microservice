package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis(uri, password string) (*redis.Client, error) {

	redisDB := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
		DB:       0,
	})

	_, err := redisDB.Ping(context.Background()).Result()
	return redisDB, err
}
