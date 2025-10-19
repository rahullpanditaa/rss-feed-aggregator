package handlers

import (
	"context"
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// HandlerUnfollow will take a feed url as cli arg, unfollow
// the feed for the currently logged in user i.e. delete
// feed_follow record
func HandlerUnfollow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.CommandArgs) != 1 {
		return cli.ErrUnfollowCommandInvalidArgs
	}

	feedURL := cmd.CommandArgs[0]
	feed, err := s.DbQueries.GetFeedsByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	err = s.DbQueries.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("%s has unfollowed from the feed: %s\n", user.Name, feedURL)
	return nil
}
