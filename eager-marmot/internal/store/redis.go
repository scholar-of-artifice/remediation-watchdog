package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Redis implements the persistence layer for processed events.
// This serves as the final sink where data is stored after being consumed from the Kafka message bus.
type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore initializes the Redis client
// It relies on the REDIS_ADDR environment variable provided in the docker-compose
func NewRedisStore(addr string) *RedisStore {
	return &RedisStore{
		Client: redis.NewClient(&redis.Options{
			Addr: addr, // example: "localhost:6379"
		}),
	}
}

// SaveEvent persists the data into Redis using the EventID as the key.
// Passing the context allows for request tracking and cancellation during the service's graceful shutdown
func (s *RedisStore) SaveEvent(ctx context.Context, id string, data interface{}) error {
	// use a zero duration for expiration
	// events will persist until manually cleared or reset
	return s.Client.Set(ctx, id, data, 0).Err()
}
