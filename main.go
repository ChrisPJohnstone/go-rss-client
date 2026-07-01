package main

import (
	"fmt"
	"log"
)

type Feed struct {
	Title       string
	Description string
	Link        string
	FeedType    string // "rss" or "atom"
	Items       []Item
}

type Item struct {
	Title       string
	Description string
	Link        string
	Author      string
	PubDate     string // could use time.Time later
	ID          string // RSS: guid, Atom: id
}

func fetchFeed(url string) (Feed, error) {
	feed, err := fetchRSSFeed(url)
	if err != nil {
		return Feed{}, err
	}
	return feed, nil
}

func main() {
	var url string = "https://feeds.bbci.co.uk/news/rss.xml"
	feed, err := fetchFeed(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
