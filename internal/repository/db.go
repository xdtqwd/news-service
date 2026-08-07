package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DbConnect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:password@localhost:5434/news_db")
}
