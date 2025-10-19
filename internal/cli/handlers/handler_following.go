package handlers

import (
	"context"
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

func HandlerFollowing(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 0 {
		return cli.ErrFollowingCommandInvalidArgs
	}

	// current user
	currentUserName := s.ApplicationState.CurrentUserName

	feedFollowsForCurrentUser, err := s.DbQueries.GetFeedFollowsForUser(context.Background(), currentUserName)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds followed by current user %s:\n", currentUserName)
	for _, feedFollow := range feedFollowsForCurrentUser {
		fmt.Printf(" - %s\n", feedFollow.FeedName)
	}
	return nil
}
