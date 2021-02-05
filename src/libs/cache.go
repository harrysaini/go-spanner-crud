package libs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache - redis related operations
type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisCacheClient - creates new redis client
func NewRedisCacheClient() *RedisCache {
	address := fmt.Sprintf("%s:%d", Conf.Redis.Host, Conf.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	ttl := Conf.Redis.TTL
	return &RedisCache{
		client: client,
		ttl:    time.Duration(ttl) * time.Second,
	}
}

// Set sets key in redis
func (cache *RedisCache) Set(ctx context.Context, key string, value []byte) error {
	return cache.client.Set(ctx, key, value, cache.ttl).Err()
}

// Get gets keys from redis
func (cache *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := cache.client.Get(ctx, key).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(val), nil
}

// Delete delete key from redis
func (cache *RedisCache) Delete(ctx context.Context, key string) error {
	return cache.client.Del(ctx, key).Err()
}
