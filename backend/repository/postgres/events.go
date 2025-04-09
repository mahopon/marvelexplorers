package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) GetEvents(ctx context.Context, offset int) interface{} {
	rows, err := pg.db.Query(ctx, "SELECT * from Events LIMIT 20 OFFSET @offset;", pgx.NamedArgs{
		"offset": offset,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer rows.Close()

	columns := rows.FieldDescriptions()
	return_objs := []interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		err := rows.Scan(values...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col.Name] = *(values[i].(*interface{}))
		}

		return_objs = append(return_objs, rowMap)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}

	return return_objs
}
