package handlers

import (
	"fmt"
	"time"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/rss"
)

// HandlerAgg -> agg command.
// Will fetch the feed from a url
// and print the struct to console
func HandlerAgg(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 1 {
		return cli.ErrAggCommandInvalidArgs
	}
	// user must provide exactly 1 argument
	// the time between requests
	time_between_reqs, err := time.ParseDuration(cmd.CommandArgs[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	// for ; ; <-ticker.C {
	// 	rss.ScrapeFeeds(s)
	// }

	for {
		rss.ScrapeFeeds(s)
		<-ticker.C
	}
}
