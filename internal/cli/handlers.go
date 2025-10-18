package cli

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// HandlerLogin will update the application state
// to reflect the user logging in
func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.CommandArgs) == 0 {
		return ErrLoginCommandInvalidArgs
	}

	usernameEntered := cmd.CommandArgs[0]
	err := s.ApplicationState.SetUser(usernameEntered)
	if err != nil {
		return err
	}
	fmt.Printf("app state has been set to given user: %s\n", usernameEntered)
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.CommandArgs) == 0 {
		return ErrRegisterCommandInvalidArgs
	}

	usernameEntered := cmd.CommandArgs[0]
	user, err := s.DbQueries.GetUser(context.Background(), usernameEntered)
	if err != nil {
		// ErrNoRows is returned by sqlc when no match found
		// in row.Scan i.e. user does not exist
		if err != sql.ErrNoRows {
			return err
		}
		// err is ErrNoRows, which means user does not exist
		// create it
		user, err = s.DbQueries.CreateUser(
			context.Background(),
			database.CreateUserParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Name:      usernameEntered,
			},
		)
		if err != nil {
			return err
		}
	} else {
		// user already exists
		fmt.Fprintln(os.Stderr, "user already exists")
		os.Exit(1)
	}

	err = s.ApplicationState.SetUser(usernameEntered)
	if err != nil {
		return err
	}
	fmt.Printf("Created user: %s\n", usernameEntered)
	fmt.Println("User ID:", user.ID)
	fmt.Println("Username:", user.Name)
	fmt.Println("Created at:", user.CreatedAt)

	return nil
}
