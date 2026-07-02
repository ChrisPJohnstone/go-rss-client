package rssclient

import "encoding/xml"

type atomFeed struct {
	Title    string     `xml:"title"`
	Subtitle string     `xml:"subtitle"`
	Link     atomLink   `xml:"link"`
	Entries  []atomItem `xml:"entry"`
}

type atomItem struct {
	Title     string   `xml:"title"`
	Link      atomLink `xml:"link"`
	ID        string   `xml:"id"`
	Published string   `xml:"published"`
	Updated   string   `xml:"updated"`
	Summary   string   `xml:"summary"`
	Content   string   `xml:"content"`
	Author    struct {
		Name string `xml:"name"`
	} `xml:"author"`
}

type atomLink struct {
	Href string `xml:"href,attr"`
}

func parseAtomFeed(body []byte) (Feed, error) {
	var atom atomFeed
	if err := xml.Unmarshal(body, &atom); err != nil {
		return Feed{}, err
	}
	feed := Feed{
		Title:       atom.Title,
		Description: atom.Subtitle,
		Link:        atom.Link.Href,
		FeedType:    "atom",
	}
	for _, entry := range atom.Entries {
		desc := entry.Summary
		if desc == "" {
			desc = entry.Content
		}
		pubDate := entry.Published
		if pubDate == "" {
			pubDate = entry.Updated
		}
		feed.Items = append(feed.Items, Item{
			Title:       entry.Title,
			Description: desc,
			Link:        entry.Link.Href,
			PubDate:     pubDate,
			Author:      entry.Author.Name,
			ID:          entry.ID,
		})
	}
	return feed, nil
}
