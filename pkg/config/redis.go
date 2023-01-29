package config

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(config *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB: 0,
	})
}

