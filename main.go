package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Bike struct {
	Brand string
	Model string
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("GO_POSTGRES"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "the database connectino filae %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// QueryRow
	var version string
	err = conn.QueryRow(context.Background(), "select version()").Scan(&version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	fmt.Println(version)

	// Query
	rows, err := conn.Query(context.Background(), "select brand, model from bikes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	// CollectRows
	bikes, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Bike])
	for _, bike := range bikes {
		fmt.Printf("The bike brand %s, model %s\n", bike.Brand, bike.Model)
	}

	// ForEachRow
	var brand, model string
	rows, _ = conn.Query(context.Background(), "select brand, model from bikes")
	_, _ = pgx.ForEachRow(rows, []any{&brand, &model}, func() error {
		fmt.Printf("%s is a model of %s\n", model, brand)
		return nil
	})
}
