package api

import (
	"database/sql"
	"log"
)

// Logger and DB conneection
type Application struct {
	DB                   *sql.DB
	Info, Error, Success *log.Logger
}
