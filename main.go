package main

import (
	"fmt"
	"log"
)

func main() {
	var url string = "https://feeds.bbci.co.uk/news/rss.xml"
	rss, err := rssFeed(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
	}
}
