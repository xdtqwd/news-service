package repository

import (
	"context"
	"news-service/internal/parser"

	"github.com/jackc/pgx/v5"
)

type Article struct {
	ID          int
	Title       string
	Link        string
	Description string
	PubDate     string
}

func GetArticle(conn *pgx.Conn) ([]Article, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT id, title, link, description, pub_date FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		err = rows.Scan(&a.ID, &a.Title, &a.Link, &a.Description, &a.PubDate)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func SaveArticle(conn *pgx.Conn, item parser.Item) error {
	_, err := conn.Exec(context.Background(),
		` INSERT INTO articles (title, link, description, pub_date)
		  VALUES ($1, $2, $3, $4)
		  ON CONFLICT (link) DO NOTHING`,
		item.Title, item.Link, item.Description, item.PubDate,
	)
	return err
}
