package main

import (
	"fmt"
	"log"
)

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
