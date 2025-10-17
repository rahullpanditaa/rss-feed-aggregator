package cli

import "fmt"

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
