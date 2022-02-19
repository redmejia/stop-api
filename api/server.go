package api

import (
	"database/sql"
	"log"
)

// Loger and DB conneection
type Application struct {
	DB                   *sql.DB
	Info, Error, Success *log.Logger
}
