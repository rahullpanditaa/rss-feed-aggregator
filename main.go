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
	// read json file, store fields in a Config struct
	appState, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "could not read config json")
		os.Exit(1)
	}

	// extract db url from config struct
	dbURL := appState.DbURL

	// create connection pool
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "could not open specified db")
		os.Exit(1)
	}
	defer db.Close()

	// get a pointer to a Queries struct, used to
	// run all sqlc generated sql query -> Go code
	dbQueries := database.New(db)

	// create State object - holds pointer to the app state,
	// and a pointer to Queries struct
	state := cli.State{
		ApplicationState: &appState,
		DbQueries:        dbQueries,
	}

	// create a Commands object
	// field CmdsRegistry is an empty map - command name -> commandHandlerFunc
	commands := cli.Commands{
		CmdsRegistry: make(map[string]func(*cli.State, cli.Command) error),
	}

	// register all CLI commands
	commands.Register("register", handlers.HandlerRegister)
	commands.Register("login", handlers.HandlerLogin)
	commands.Register("reset", handlers.HandlerReset)
	commands.Register("users", handlers.HandlerUsers)
	commands.Register("agg", handlers.HandlerAgg)
	commands.Register("feeds", handlers.HandlerFeeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(handlers.HandlerFollow))
	commands.Register("following", cli.MiddlewareLoggedIn(handlers.HandlerFollowing))
	commands.Register("addfeed", cli.MiddlewareLoggedIn(handlers.HandlerAddFeed))
	commands.Register("unfollow", cli.MiddlewareLoggedIn(handlers.HandlerUnfollow))
	commands.Register("browse", cli.MiddlewareLoggedIn(handlers.HandlerBrowse))

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
