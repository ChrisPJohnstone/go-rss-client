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
	FeedType    FeedType
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

func parseFeedType(body []byte) FeedType {
	if bytes.Contains(body[:200], []byte(`<rss`)) {
		return FeedTypeRSS
	}
	if bytes.Contains(body[:200], []byte(`<feed`)) {
		return FeedTypeAtom
	}
	return FeedTypeUnknown
}

func FetchFeed(url string) (Feed, error) {
	body, err := fetch(url)
	if err != nil {
		return Feed{}, err
	}
	switch parseFeedType(body) {
	case FeedTypeRSS:
		return parseRSSFeed(body)
	case FeedTypeAtom:
		return parseAtomFeed(body)
	default:
		return Feed{}, fmt.Errorf("unknown feed type")
	}
}
