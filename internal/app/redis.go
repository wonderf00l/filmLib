package app

import (
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
	"github.com/wonderf00l/filmLib/internal/configs"
)

func NewRedisClient(ctx context.Context, cfg configs.RedisConfig) (*redis.Client, error) {
	redisCl := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
	})

	status := redisCl.Ping(ctx)
	if status.Err() != nil {
		redisCl.Close()
		return nil, fmt.Errorf("new redis client: %w", status.Err())
	}
	return redisCl, nil
}
