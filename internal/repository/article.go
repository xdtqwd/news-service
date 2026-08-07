package repository

import (
	"context"
	"news-service/internal/parser"

	"github.com/jackc/pgx/v5"
)

func SaveArticle(conn *pgx.Conn, item parser.Item) error {
	_, err := conn.Exec(context.Background(),
		` INSERT INTO articles (title, link, description, pub_date)
		  VALUES ($1, $2, $3, $4)
		  ON CONFLICT (link) DO NOTHING`,
		item.Title, item.Link, item.Description, item.PubDate,
	)
	return err
}
