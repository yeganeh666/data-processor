package redisext

import "github.com/go-redis/redis/v8"

func NewClient(address, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis",
	})
}
