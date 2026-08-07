package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Migrate(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS articles(
	id 			SERIAL PRIMARY KEY,
	title		TEXT NOT NULL,
	link		TEXT UNIQUE NOT NULL,
	description	TEXT,
	pub_date	TEXT,
	created_at	TIMESTAMP DEFAULT NOW()
	
	
	
	
	)`)
	return err
}

func DbConnect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:password@localhost:5434/news_db")
}
