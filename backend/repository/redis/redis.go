package redis

import (
	redislib "github.com/mahopon/gobackend/redis"
	"sync"
)

type RedisRepo[T any] struct {
	client redislib.RedisClient
}

var (
	client redislib.RedisClient
	once   sync.Once
)

func NewRedisRepo[T any]() *RedisRepo[T] {
	return &RedisRepo[T]{client: getClient()}
}

func getClient() redislib.RedisClient {
	once.Do(func() {
		client = redislib.GetClient()
	})
	return client
}
