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

func fetch(url string, headers http.Header) ([]byte, http.Header, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	for key, values := range headers {
		for _, v := range values {
			req.Header.Add(key, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, resp.Header, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.Header, err
	}
	return body, resp.Header, nil
}

func parseFeedType(body []byte) FeedType {
	n := min(len(body), 100)
	prefix := body[:n]
	if bytes.Contains(prefix, []byte(`<rss`)) {
		return FeedTypeRSS
	}
	if bytes.Contains(prefix, []byte(`<feed`)) {
		return FeedTypeAtom
	}
	return FeedTypeUnknown
}

func FetchFeed(url string, headers http.Header) (Feed, http.Header, error) {
	var body []byte
	var header http.Header
	var err error
	body, header, err = fetch(url, headers)
	if err != nil {
		return Feed{}, header, err
	}
	var feed Feed
	switch parseFeedType(body) {
	case FeedTypeRSS:
		feed, err = parseRSSFeed(body)
	case FeedTypeAtom:
		feed, err = parseAtomFeed(body)
	default:
		return Feed{}, header, fmt.Errorf("unknown feed type")
	}
	if err != nil {
		return Feed{}, header, err
	}
	return feed, header, nil
}
