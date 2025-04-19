package redis

import (
	redis "github.com/mahopon/gobackend/redis"
)

type CharacterRepoRedis struct {
	client redis.RedisClient
}

func NewCharacterRepoRedis() *CharacterRepoRedis {
	return &CharacterRepoRedis{
		client: redis.GetClient(),
	}
}
