package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// HandlerRegister is called when the register command is
// entered in the cli. Create a new user in the db and
// updates the config struct to hold the username
func HandlerRegister(s *cli.State, cmd cli.Command) error {
	if len(cmd.CommandArgs) == 0 {
		return cli.ErrRegisterCommandInvalidArgs
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
