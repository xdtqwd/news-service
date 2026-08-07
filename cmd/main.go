package main

import (
	"fmt"

	"net/http"
	"news-service/internal/parser"
)

func main() {
	items, err := parser.Parse("https://feeds.bbci.co.uk/news/rss.xml")
	fmt.Println("err:", err)
	fmt.Println("items count:", len(items))
	for _, item := range items {
		fmt.Println(item.Title)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	})

	http.ListenAndServe(":8080", nil)
}
