package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sync"
)

type PostgresRepo[T any] struct {
	db *pgxpool.Pool
}

var (
	dbPool *pgxpool.Pool
	dbOnce sync.Once
)

func NewPGRepo[T any]() *PostgresRepo[T] {
	return &PostgresRepo[T]{db: getDBPool()}
}

func getDBPool() *pgxpool.Pool {
	dbOnce.Do(func() {
		ctx := context.Background()
		var err error
		dbPool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Printf("unable to create connection pool: %v\n", err)
			return
		}
	})
	return dbPool
}

func (pg *PostgresRepo[T]) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *PostgresRepo[T]) Close() {
	pg.db.Close()
}
