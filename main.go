package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli/handlers"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/config"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

func main() {
	// read json file, store fields in a struct
	appState, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "could not read config json")
		os.Exit(1)
	}

	dbURL := appState.DbURL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "could not open specified db")
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	state := cli.State{
		ApplicationState: &appState,
		DbQueries:        dbQueries,
	}

	commands := cli.Commands{
		CmdsRegistry: make(map[string]func(*cli.State, cli.Command) error),
	}
	commands.Register("register", handlers.HandlerRegister)
	commands.Register("login", handlers.HandlerLogin)
	commands.Register("reset", handlers.HandlerReset)
	commands.Register("users", handlers.HandlerUsers)
	commands.Register("agg", handlers.HandlerAgg)
	commands.Register("addfeed", handlers.HandlerAddFeed)
	commands.Register("feeds", handlers.HandlerFeeds)
	commands.Register("follow", handlers.HandlerFollow)
	commands.Register("following", handlers.HandlerFollowing)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "too few command-line arguments")
		os.Exit(1)
	}
	cmdNameEntered := os.Args[1]
	cmdArgsEntered := os.Args[2:]

	cmd := cli.Command{
		CommandName: cmdNameEntered,
		CommandArgs: cmdArgsEntered,
	}

	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
