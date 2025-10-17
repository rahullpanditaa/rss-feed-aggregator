package cli

import "errors"

var (
	ErrLoginCommandInvalidArgs = errors.New("usage: login <username>")
)
