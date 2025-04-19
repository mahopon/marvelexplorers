package repository

import (
	"context"
	"time"
)

type DBRepo[T any] interface {
	Get(ctx context.Context, table string, offset int) ([]T, error)
	Search(ctx context.Context, table string, searchString string) ([]T, error)
}

type CacheRepo[T any] interface {
	Get(ctx context.Context, table string, searchString string) (string, error)
	Insert(ctx context.Context, table string, key string, value any, ttl time.Duration) error
}
