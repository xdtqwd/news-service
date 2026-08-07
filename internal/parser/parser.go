package parser

import (
	"encoding/xml"
	"net/http"
)

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}

func Parse(url string) ([]Item, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	return rss.Channel.Items, nil
}
