package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func (pg *PostgresRepo[T]) Get(ctx context.Context, table string, offset int) ([]T, error) {
	rows, err := pg.db.Query(ctx, "SELECT * from "+table+" LIMIT 20 OFFSET @offset;", pgx.NamedArgs{
		"offset": offset,
	})
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		return nil, err
	}
	return records, nil
}

// func (pg *Postgres) InsertCharacters(ctx context.Context)

func (pg *PostgresRepo[T]) Search(ctx context.Context, table string, searchString string) ([]T, error) {
	rows, err := pg.db.Query(ctx, "SELECT * FROM "+table+" WHERE name ilike '%' || @search || '%'", pgx.NamedArgs{
		"search": searchString,
	})
	if err != nil {
		log.Printf("SearchCharacter: Error caused by search character query: %v", err)
		return nil, err
	}

	characters, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		log.Printf("SearchCharacter: Error unpacking character values: %v", err)
		return nil, err
	}
	return characters, nil
}
