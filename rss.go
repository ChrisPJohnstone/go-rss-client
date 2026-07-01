package main

import "encoding/xml"

type rssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []rssItem `xml:"item"`
	} `xml:"channel"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Author      string `xml:"author"`
	GUID        string `xml:"guid"`
}

func parseRSSFeed(body []byte) (Feed, error) {
	var rss rssFeed
	if err := xml.Unmarshal(body, &rss); err != nil {
		return Feed{}, err
	}
	feed := Feed{
		Title:       rss.Channel.Title,
		Description: rss.Channel.Description,
		Link:        rss.Channel.Link,
		FeedType:    "rss",
	}
	for _, item := range rss.Channel.Items {
		feed.Items = append(feed.Items, Item{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PubDate:     item.PubDate,
			Author:      item.Author,
			ID:          item.GUID,
		})
	}
	return feed, nil
}
