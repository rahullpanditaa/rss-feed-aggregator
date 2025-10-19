package handlers

import (
	"context"
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

// HandlerFeeds lists out all feeds created by current user
func HandlerFeeds(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 0 {
		return cli.ErrFeedsCommandInvalidArgs
	}

	allFeeds, err := s.DbQueries.GetAllFeedsWithCreatorUsername(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range allFeeds {
		fmt.Printf("Feed name: %s, feed url: %s, created by %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil

}
