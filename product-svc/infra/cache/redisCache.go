package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/slg"
	"time"
)

type RedisCache interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}

type redisCache struct {
	client *redis.Client
}

func (r *redisCache) Get(ctx context.Context, key string, dest any) error {
	cached, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return redis.Nil
	}
	if err != nil {
		slg.Logger.Warn("redis get error", "key", key, "err", err.Error())
		return err
	}

	if err := json.Unmarshal([]byte(cached), dest); err != nil {
		slg.Logger.Warn("failed to unmarshal cached data", "key", key, "err", err.Error())
		return err
	}

	slg.Logger.Info("cache hit", "key", key)
	return nil
}

func (r *redisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		slg.Logger.Warn("failed to marshal cache data", "key", key, "err", err.Error())
		return err
	}

	if err := r.client.Set(ctx, key, data, ttl).Err(); err != nil {
		slg.Logger.Warn("redis set error", "key", key, "err", err.Error())
		return err
	}

	return nil
}

func (r *redisCache) Delete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	if err := r.client.Del(ctx, keys...).Err(); err != nil {
		slg.Logger.Warn("failed to delete cache keys", "keys", keys, "err", err.Error())
		return err
	}
	return nil
}

func NewRedisCache(client *redis.Client) RedisCache {
	return &redisCache{
		client: client,
	}
}
