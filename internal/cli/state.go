package cli

import (
	"github.com/rahullpanditaa/rssfeedaggregator/internal/config"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

// State struct holds a pointer to a Config struct
// which represents the current state of the application
// in the form of unmarshaled config json
// and a pointer to a Queries struct which has all sqlc
// generated query methods
type State struct {
	ApplicationState *config.Config
	DbQueries        *database.Queries
}
