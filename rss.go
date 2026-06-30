package main

import (
	"encoding/xml"
	"fmt"
	"io"
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

func rssFeed(url string) (RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return RSS{}, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSS{}, err
	}
	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return RSS{}, err
	}
	return rss, nil
}
