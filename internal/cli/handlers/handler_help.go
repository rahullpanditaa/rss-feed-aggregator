package handlers

import (
	"fmt"

	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
)

func HandlerHelp(s *cli.State, c cli.Command) error {
	fmt.Println("Available commands:")
	fmt.Println()

	fmt.Println("  register <username>       - Create a new user")
	fmt.Println("  login <username>          - Log in as an existing user")
	fmt.Println("  users                     - List all users")
	fmt.Println("  reset                     - Reset database, delete all users")
	fmt.Println("  feeds                     - List all available feeds")
	fmt.Println("  addfeed <name> <url>      - Add a new feed (and follow it.)")
	fmt.Println("  follow <url>              - Follow an existing feed")
	fmt.Println("  unfollow <url>            - Unfollow a feed")
	fmt.Println("  following                 - Show feeds you follow")
	fmt.Println("  browse <limit>            - Show recent posts")
	fmt.Println("  agg <duration>            - Start aggregator loop")
	fmt.Println("  help                      - Show this help menu")

	return nil
}
