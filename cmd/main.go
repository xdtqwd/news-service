package main

import (
	"context"
	"fmt"

	"net/http"
	"news-service/internal/parser"
	"news-service/internal/repository"
)

func main() {
	items, err := parser.Parse("https://feeds.bbci.co.uk/news/rss.xml")
	fmt.Println("err:", err)
	fmt.Println("items count:", len(items))
	for _, item := range items {
		fmt.Println(item.Title)
	}
	conn, err := repository.DbConnect(context.Background())
	if err != nil {
		fmt.Println("Ошибка БД!", err)
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("БД подлючена!")

	err = repository.Migrate(conn)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	fmt.Println("Миграция выполнена!")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	})

	http.ListenAndServe(":8080", nil)
}
