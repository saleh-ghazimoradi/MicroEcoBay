package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr, pw string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	
	return client, nil
}
