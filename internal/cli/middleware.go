package cli

import (
	"context"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	f := func(s *State, c Command) error {
		currentUsername := s.ApplicationState.CurrentUserName
		user, err := s.DbQueries.GetUser(context.Background(), currentUsername)
		if err != nil {
			return err
		}
		if user.Name != currentUsername {
			return ErrUserDoesNotExist
		}
		return handler(s, c, user)
	}
	return f
}
