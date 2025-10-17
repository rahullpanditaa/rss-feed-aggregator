package cli

import "fmt"

type Command struct {
	CommandName string
	CommandArgs []string
}

type Commands struct {
	CmdsRegistry map[string]func(*State, Command) error
}

// register will register a new handler function for
// a command name.
func (c *Commands) Register(name string, f func(*State, Command) error) {
	_, exists := c.CmdsRegistry[name]
	if exists {
		fmt.Printf("given command %s already exists", name)
		return
	}
	c.CmdsRegistry[name] = f
}

// run executes the given cli command if it exists with
// the provided state.
func (c *Commands) Run(s *State, cmd Command) error {
	cmdFunc, exists := c.CmdsRegistry[cmd.CommandName]
	if !exists {
		return ErrCommandDoesNotExist
	}

	err := cmdFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
