package cli

import "errors"

var (
	ErrLoginCommandInvalidArgs     = errors.New("usage: login <username>")
	ErrRegisterCommandInvalidArgs  = errors.New("usage: register <username>")
	ErrCommandDoesNotExist         = errors.New("cannot run command as it doesn't exist")
	ErrUserDoesNotExist            = errors.New("user does not exist")
	ErrResetCommandInvalidArgs     = errors.New("reset does not take any arguments")
	ErrUsersCommandInvalidArgs     = errors.New("users does not take any arguments")
	ErrAddFeedCommandInvalidArgs   = errors.New("usage: addfeed <feed_name> <feed_url>")
	ErrFeedsCommandInvalidArgs     = errors.New("feeds does not take any arguments")
	ErrFollowCommandInvalidArgs    = errors.New("usage: follow <feed_url>")
	ErrFollowingCommandInvalidArgs = errors.New("following does not take any arguments")
)
