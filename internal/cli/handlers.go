package cli

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/rss"
)

// HandlerLogin will update the application state
// to reflect the user logging in
func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.CommandArgs) == 0 {
		return ErrLoginCommandInvalidArgs
	}

	usernameEntered := cmd.CommandArgs[0]
	_, err := s.DbQueries.GetUser(context.Background(), usernameEntered)
	if err != nil {
		if err == sql.ErrNoRows {
			// user does not exist
			fmt.Fprintln(os.Stderr, ErrUserDoesNotExist)
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

// HandlerRegister is called when the register command is
// entered in the cli. Create a new user in the db and
// updates the config struct to hold the username
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

// HandlerReset resets the users table
func HandlerReset(s *State, cmd Command) error {
	if len(cmd.CommandArgs) != 0 {
		return ErrResetCommandInvalidArgs
	}
	err := s.DbQueries.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not reset database:", err)
		os.Exit(1)
	}
	fmt.Println("Database reset, deleted all users")
	return nil
}

// HandlerUsers will print all the users to the console,
// along with specifying which user is currently logged-in
func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.CommandArgs) != 0 {
		return ErrUsersCommandInvalidArgs
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

// HandlerAgg -> agg command.
// Will fetch the feed from a url (single, hardcoded for now)
// and print the struct to console
func HandlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feedStruct, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(feedStruct, "", "  ")
	fmt.Println(string(b))
	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	currentUser := s.ApplicationState.CurrentUserName
	user, err := s.DbQueries.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}
	currentUserId := user.ID

	if len(cmd.CommandArgs) != 2 {
		return ErrAddFeedCommandInvalidArgs
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
	fmt.Printf("Feed ID: %v\n", feed.ID)
	fmt.Printf("Feed created at: %v\n", feed.CreatedAt)
	fmt.Printf("Feed updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("Feed name: %s\n", feed.Name)
	fmt.Printf("Feed urp: %s\n", feed.Url)
	fmt.Printf("User connected to this feed: %v\n", feed.UserID)

	return nil

}
