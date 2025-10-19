package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

// HandlerUsers will print all the users to the console,
// along with specifying which user is currently logged-in
func HandlerUsers(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 0 {
		return cli.ErrUsersCommandInvalidArgs
	}
	allUsers, err := s.DbQueries.GetUsers(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not retrieve all users:", err)
		os.Exit(1)
	}
	if len(allUsers) == 0 {
		fmt.Println("no users found")
		return nil
	}
	for _, user := range allUsers {
		if user.Name == s.ApplicationState.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
