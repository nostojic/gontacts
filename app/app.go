package app

import "github.com/jackc/pgx/v5/pgxpool"

type App struct {
	Db *pgxpool.Pool
}