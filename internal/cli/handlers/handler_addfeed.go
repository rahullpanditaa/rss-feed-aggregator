package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// HandlerAddFeed creates a new feed, given a feed name and url as command line arguments.
// It also creates a new feed_follow record afterwards
func HandlerAddFeed(s *cli.State, cmd cli.Command) error {
	currentUser := s.ApplicationState.CurrentUserName
	user, err := s.DbQueries.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}
	currentUserId := user.ID

	if len(cmd.CommandArgs) != 2 {
		return cli.ErrAddFeedCommandInvalidArgs
	}
	feedName := cmd.CommandArgs[0]
	feedUrl := cmd.CommandArgs[1]

	feed, err := s.DbQueries.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      feedName,
			Url:       feedUrl,
			UserID:    currentUserId,
		},
	)
	if err != nil {
		return err
	}

	// create a feed follow record after adding a feed
	_, err = s.DbQueries.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    feed.UserID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Feed ID: %v\n", feed.ID)
	fmt.Printf("Feed created at: %v\n", feed.CreatedAt)
	fmt.Printf("Feed updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("Feed name: %s\n", feed.Name)
	fmt.Printf("Feed url: %s\n", feed.Url)
	fmt.Printf("User connected to this feed: %v\n", feed.UserID)

	return nil
}
