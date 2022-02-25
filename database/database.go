package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// PostgreQL db
type DB struct {
	Db *sql.DB
}

const (
	openConns = 10
	idleConns = 3
	lifeTime  = 60 * time.Second
)

// connection
func Connection() (*DB, error) {

	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("DBSSLMODE"),
	)

	var db DB
	var err error

	db.Db, err = sql.Open("pgx", connDB)
	if err != nil {
		return nil, err
	}

	db.Db.SetMaxOpenConns(openConns)
	db.Db.SetMaxIdleConns(idleConns)
	db.Db.SetConnMaxLifetime(lifeTime)

	return &db, nil
}

// test method
func (d *DB) Get() {
	fmt.Println("Getting Gophers")
}

// connection test
func ConnectionPing(db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
