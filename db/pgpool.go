package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDb() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("PG_URL"))

	if err != nil {
		return pool, fmt.Errorf("unable to connect to db: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close() // Clean up
		return pool, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}