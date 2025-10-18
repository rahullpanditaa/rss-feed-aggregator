package cli

import "errors"

var (
	ErrLoginCommandInvalidArgs    = errors.New("usage: login <username>")
	ErrRegisterCommandInvalidArgs = errors.New("usage: register <username>")
	ErrCommandDoesNotExist        = errors.New("cannot run command as it doesn't exist")
	ErrUserDoesNotExist           = errors.New("user does not exist")
)
