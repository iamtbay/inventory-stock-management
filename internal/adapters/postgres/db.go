package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(connStr string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), config)
}
