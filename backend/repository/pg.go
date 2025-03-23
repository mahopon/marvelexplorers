package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connString string) (*postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &postgres{db}
	})

	return pgInstance, nil
}

func GetPG() *postgres {
	return pgInstance
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
