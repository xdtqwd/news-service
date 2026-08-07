package handler

import (
	"encoding/json"
	"net/http"
	"news-service/internal/repository"

	"github.com/jackc/pgx/v5"
)

func GetArticles(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := repository.GetArticle(conn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}
}
