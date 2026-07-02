package rssclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func feedType(body []byte) string {
	if bytes.Contains(body[:200], []byte(`<rss`)) {
		return "rss"
	}
	if bytes.Contains(body[:200], []byte(`<feed`)) {
		return "atom"
	}
	return ""
}

func FetchFeed(url string) (Feed, error) {
	body, err := fetch(url)
	if err != nil {
		return Feed{}, err
	}
	switch feedType(body) {
	case "rss":
		return parseRSSFeed(body)
	case "atom":
		return parseAtomFeed(body)
	default:
		return Feed{}, fmt.Errorf("unknown feed type")
	}
}
