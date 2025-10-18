package cli

import (
	"fmt"
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

	return nil
}
