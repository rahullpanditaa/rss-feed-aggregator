package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// HandlerFollow creates a new feed_follows record for
// the current user. It also prints the name of the feed
// and the current user
func HandlerFollow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.CommandArgs) != 1 {
		return cli.ErrFollowCommandInvalidArgs
	}
	feedUrl := cmd.CommandArgs[0]

	feed, err := s.DbQueries.GetFeedsByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	// create new feed follow record
	feed_follow_row, err := s.DbQueries.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Feed name: %s\n", feed_follow_row.FeedName)
	fmt.Printf("Current user: %s\n", feed_follow_row.UserName)

	return nil

}
