package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/rss"
)

// HandlerAgg -> agg command.
// Will fetch the feed from a url (single, hardcoded for now)
// and print the struct to console
func HandlerAgg(s *cli.State, cmd cli.Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feedStruct, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(feedStruct, "", "  ")
	fmt.Println(string(b))
	return nil
}
