package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Cache Struct
type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

// NewCache inits a redis cache
func NewCache(redisAddress string, password ...string) (*Cache, error) {
	ttl := 24 * time.Hour

	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

// Add key value pair to cache
func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

// Get key value pair from cache
func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
