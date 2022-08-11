package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type storage struct {
	rdb *redis.Client
}

func NewStorage(rdb *redis.Client) *storage {
	return &storage{rdb: rdb}
}

func (s *storage) Set(modelId int) error {

	_, err := s.rdb.Set(context.Background(), string(modelId), true, 0).Result()
	if err != nil {
		return err
	}

	return nil
}
