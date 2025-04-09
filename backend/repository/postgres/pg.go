package postgres

import (
	"context"
	"fmt"
	"sync"
	repo "tcy/marvelexplorers/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

var _ repo.CharacterRepo = (*Postgres)(nil)

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func GetPG() *Postgres {
	if pgInstance == nil {
		return nil
	} else {
		return pgInstance
	}
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
