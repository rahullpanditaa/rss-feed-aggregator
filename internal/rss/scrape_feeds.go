package rss

import (
	"context"
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

func ScrapeFeeds(s *cli.State) {
	// get the next feed to fetch
	feed, err := s.DbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get feed to fetch: %v\n", err)
		os.Exit(1)
	}

	// mark it as fetched
	err = s.DbQueries.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not mark feed as fetched: %v\n", err)
		os.Exit(1)
	}

	// fetch feed using feed url
	feedStruct, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "coud not fetch feed: %v\n", err)
		os.Exit(1)
	}

	// iterate over ITEMS in feed
	for _, item := range feedStruct.Channel.Item {
		fmt.Printf("Item title: %s\n", item.Title)
	}
}
