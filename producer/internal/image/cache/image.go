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

func (s *storage) CheckInCache(id int64) bool {

	_, err := s.rdb.Get(context.Background(), strconv.Itoa(int(id))).Result()
	if err != nil {
		return false
	}

	return true
}
