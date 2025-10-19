package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

// HandlerLogin will update the application state
// to reflect the user logging in
func HandlerLogin(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) == 0 {
		return cli.ErrLoginCommandInvalidArgs
	}

	usernameEntered := cmd.CommandArgs[0]
	_, err := s.DbQueries.GetUser(context.Background(), usernameEntered)
	if err != nil {
		if err == sql.ErrNoRows {
			// user does not exist
			fmt.Fprintln(os.Stderr, cli.ErrUserDoesNotExist)
			os.Exit(1)
		}
		return err
	}
	err = s.ApplicationState.SetUser(usernameEntered)
	if err != nil {
		return err
	}
	fmt.Printf("app state has been set to given user: %s\n", usernameEntered)
	return nil
}
