package handlers

import (
	"context"
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

func HandlerFollowing(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.CommandArgs) != 0 {
		return cli.ErrFollowingCommandInvalidArgs
	}

	feedFollowsForCurrentUser, err := s.DbQueries.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds followed by current user %s:\n", user.Name)
	for _, feedFollow := range feedFollowsForCurrentUser {
		fmt.Printf(" - %s\n", feedFollow.FeedName)
	}
	return nil
}
