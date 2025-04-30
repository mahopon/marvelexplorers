package redis

import (
	"context"
	"strings"
	"time"
)

func (repo *RedisRepo[T]) Get(ctx context.Context, table string, searchString string) (string, error) {
	var builder strings.Builder
	builder.WriteString(table)
	builder.WriteString("|")
	builder.WriteString(searchString)
	result, err := repo.client.Get(builder.String())
	if err != nil {
		return "", err
	}
	return result, nil
}

func (repo *RedisRepo[T]) Insert(ctx context.Context, table string, key string, value any, ttl time.Duration) error {
	var builder strings.Builder
	builder.WriteString(table)
	builder.WriteString("|")
	builder.WriteString(key)
	err := repo.client.Set(builder.String(), value, ttl)
	if err != nil {
		return err
	}
	return nil
}
