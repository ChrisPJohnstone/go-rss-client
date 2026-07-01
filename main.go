package main

import (
	"fmt"
	"log"
)

func main() {
	var url string = "https://github.com/webex/webex-js-sdk/releases.atom"
	feed, err := fetchFeed(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
