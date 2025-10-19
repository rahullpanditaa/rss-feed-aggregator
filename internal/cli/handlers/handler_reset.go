package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

// HandlerReset resets the users table
func HandlerReset(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) != 0 {
		return cli.ErrResetCommandInvalidArgs
	}
	err := s.DbQueries.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not reset database:", err)
		os.Exit(1)
	}
	fmt.Println("Database reset, deleted all users")
	return nil
}
