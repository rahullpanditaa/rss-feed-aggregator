package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

func HandlerBrowse(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.CommandArgs) > 1 {
		return cli.ErrBrowseCommandInvalidArgs
	}

	limit := 2
	if len(cmd.CommandArgs) != 0 {
		limit, _ = strconv.Atoi(cmd.CommandArgs[0])
	}

	posts, err := s.DbQueries.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			Name:  s.ApplicationState.CurrentUserName,
			Limit: int32(limit),
		},
	)
	if err != nil {
		return err
	}
	// fmt.Printf("Posts for user %s, (DISPLAY MAX %d posts)\n", user.Name, limit)
	for _, post := range posts {
		fmt.Printf("Id: %v\n", post.ID)
		fmt.Printf("Title: %s\n", post.Description.String)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Println()
	}

	return nil
}
