package redis

import (
	"sync"

	redislib "github.com/mahopon/gobackend/redis"
)

type RedisRepo[T any] struct {
	Client redislib.RedisClient
}

var (
	redisRepoInstance *RedisRepo[any]
	once              sync.Once
)

func GetRedisRepo() *RedisRepo[any] {
	once.Do(func() {
		redisRepoInstance = &RedisRepo[any]{
			Client: redislib.GetClient(),
		}
	})
	return redisRepoInstance
}
