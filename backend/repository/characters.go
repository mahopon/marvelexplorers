package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (pg *postgres) GetCharacters(ctx context.Context, offset int) interface{} {
	rows, err := pg.db.Query(ctx, "SELECT * from Characters LIMIT 20 OFFSET @offset;", pgx.NamedArgs{
		"offset": offset,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer rows.Close()

	// Get column count from the rows
	columns := rows.FieldDescriptions()

	return_objs := []interface{}{}

	// Prepare a slice to hold the values for each row
	for rows.Next() {
		// Create a slice of empty interfaces to hold the values of each column
		values := make([]interface{}, len(columns))
		for i := range values {
			// In go, new() CREATES A POINTER, so we need to deference it using *
			values[i] = new(interface{})
		}

		// Scan the values from the row into the slice
		// Scan sets the value AT THE MEMORY LOCATION
		err := rows.Scan(values...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}

		// Create a map to hold the column names and their corresponding values
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			// Dereference the pointer to get the value
			rowMap[col.Name] = *(values[i].(*interface{}))
		}

		// Print the values for the current row
		// for i, col := range columns {
		// 	// Print the column name and its corresponding value
		// 	fmt.Printf("%s: %v\n", col.Name, *(values[i].(*interface{})))
		// }
		// fmt.Println()
		return_objs = append(return_objs, rowMap)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}

	return return_objs
}
