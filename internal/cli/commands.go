package cli

import "fmt"

// Represents a single cli command.
// CommandName - name of the cli command
// CommandArgs - slice of arguments for the command
type Command struct {
	CommandName string
	CommandArgs []string
}

// Commands struct holds a map that acts as a registry
// of valid cli commands for the application
// map key - name of command,
// value - corresponding handler function
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
