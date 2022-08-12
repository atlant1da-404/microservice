package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type storage struct {
	rdb *redis.Client
}

func NewStorage(rdb *redis.Client) *storage {
	return &storage{rdb: rdb}
}

func (s *storage) Set(modelId int) error {
	return s.rdb.Set(context.Background(), strconv.Itoa(modelId), true, -1).Err()
}
