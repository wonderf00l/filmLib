package configs

import "os"

type RedisConfig struct {
	Addr     string
	Password string
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}
