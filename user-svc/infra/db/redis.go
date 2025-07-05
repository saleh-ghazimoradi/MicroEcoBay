package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedis(addr, pw string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	if err := client.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("redis connect fail: %w", err.Err())
	}

	return nil, nil
}
