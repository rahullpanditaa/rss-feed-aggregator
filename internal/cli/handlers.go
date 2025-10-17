package cli

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArgs) == 0 {
		return ErrLoginCommandInvalidArgs
	}

	usernameEntered := cmd.commandArgs[0]
	err := s.applicationState.SetUser(usernameEntered)
	if err != nil {
		return err
	}
	fmt.Printf("app state has been set to given user: %s\n", usernameEntered)
	return nil
}
