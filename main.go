package main

import (
	"fmt"
	"os"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/config"
)

func main() {
	// read json file, store fields in a struct
	appState, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	state := cli.State{ApplicationState: &appState}
	commands := cli.Commands{CmdsRegistry: make(map[string]func(*cli.State, cli.Command) error)}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "too few command-line arguments")
		os.Exit(-1)
	}

	cmdNameEntered := os.Args[1]
	cmdArgsEntered := os.Args[2:]

	cmd := cli.Command{CommandName: cmdNameEntered,
		CommandArgs: cmdArgsEntered}

	commands.Register(cmdNameEntered, cli.HandlerLogin)
	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
