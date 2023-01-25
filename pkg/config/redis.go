package config

import (
	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	env := LoadEnv(".")
	return redis.NewClient(&redis.Options{
		Addr: env.RedisHost + ":" + env.RedisPort,
		Password: env.RedisPassword,
		DB: 0,
	})
}
