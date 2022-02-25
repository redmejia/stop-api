package api

import (
	"log"

	"github.com/redmejia/stop/database"
)

// Logger and DB conneection
type Application struct {
	DB                   *database.DB
	Info, Error, Success *log.Logger
}
