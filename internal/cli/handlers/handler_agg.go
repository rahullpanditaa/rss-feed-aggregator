package handlers

import (
	"fmt"
	"time"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/rss"
)

// HandlerAgg -> agg command.
// Will fetch the feed from a url (single, hardcoded for now)
// and print the struct to console
func HandlerAgg(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 1 {
		return cli.ErrAggCommandInvalidArgs
	}

	time_between_reqs, err := time.ParseDuration(cmd.CommandArgs[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		rss.ScrapeFeeds(s)
	}

	// url := "https://www.wagslane.dev/index.xml"
	// feedStruct, err := rss.FetchFeed(context.Background(), url)
	// if err != nil {
	// 	return err
	// }

	// b, _ := json.MarshalIndent(feedStruct, "", "  ")
	// fmt.Println(string(b))
	// return nil
}
