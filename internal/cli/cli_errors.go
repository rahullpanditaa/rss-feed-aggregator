package cli

import "errors"

var (
	ErrLoginCommandInvalidArgs = errors.New("usage: login <username>")
	ErrCommandDoesNotExist     = errors.New("cannot run command as it doesn't exist")
)
