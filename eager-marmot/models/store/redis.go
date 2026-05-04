package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	return &RedisStore{
		Client: redis.NewClient(&redis.Options{
			Addr: addr, // example: "localhost:6379"
		}),
	}
}

func (s *RedisStore) SaveEvent(ctx context.Context, id string, data interface{}) error {
	// right now... just write directly to verify connection
	// TODO: change later
	return s.Client.Set(ctx, id, data, 0).Err()
}
