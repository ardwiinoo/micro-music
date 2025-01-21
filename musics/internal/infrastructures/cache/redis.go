package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ardwiinoo/micro-music/musics/internal/applications/cache"

)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) cache.CacheManager {
	return &redisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}

// Delete implements cache.CacheManager.
func (r *redisCache) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

// Get implements cache.CacheManager.
func (r *redisCache) Get(key string) (interface{}, error) {
	return r.client.Get(context.Background(), key).Result()
}

// Set implements cache.CacheManager.
func (r *redisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}