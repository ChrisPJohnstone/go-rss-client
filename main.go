package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RSS struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Author      string `xml:"author"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {
	var url string = "https://feeds.bbci.co.uk/news/rss.xml"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		log.Fatal(err)
	}
	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
	}
}
