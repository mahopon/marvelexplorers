package redis

import (
	"context"
	"log"
	"strings"
	"time"
)

func (repo *RedisRepo[T]) Get(ctx context.Context, table string, searchString string) (string, error) {
	var builder strings.Builder
	builder.WriteString(table)
	builder.WriteString("|")
	builder.WriteString(searchString)
	result, err := repo.Client.Get(builder.String())
	if err != nil {
		log.Printf("Error getting character, %s: %v", searchString, err)
		return "", err
	}
	return result, nil
}

func (repo *RedisRepo[T]) Insert(ctx context.Context, table string, key string, value any, ttl time.Duration) error {
	var builder strings.Builder
	builder.WriteString(table)
	builder.WriteString("|")
	builder.WriteString(key)
	err := repo.Client.Set(builder.String(), value, ttl)
	if err != nil {
		log.Printf("Could not set in Redis: %v", err)
		return err
	}
	return nil
}
