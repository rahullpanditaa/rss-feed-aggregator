package cli

import "fmt"

type command struct {
	commandName string
	commandArgs []string
}

type commands struct {
	cmdNamesAndHandlers map[string]func(*state, command) error
}

// register will register a new handler function for
// a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	_, exists := c.cmdNamesAndHandlers[name]
	if exists {
		fmt.Printf("given command %s already exists", name)
		return
	}
	c.cmdNamesAndHandlers[name] = f
}

// run executes the given cli command if it exists with
// the provided state.
func (c *commands) run(s *state, cmd command) error {
	cmdFunc, exists := c.cmdNamesAndHandlers[cmd.commandName]
	if !exists {
		return ErrCommandDoesNotExist
	}

	err := cmdFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
