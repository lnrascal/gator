package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (rss *RSSFeed) UnescapeString() {
	html.UnescapeString(rss.Channel.Title)
	html.UnescapeString(rss.Channel.Description)

	for _, item := range rss.Channel.Item {
		item.UnescapeString()
	}
}

func (rss *RSSItem) UnescapeString() {
	html.UnescapeString(rss.Title)
	html.UnescapeString(rss.Description)
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rss RSSFeed

	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return &RSSFeed{}, err
	}

	rss.UnescapeString()

	return &rss, nil
}
