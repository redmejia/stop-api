package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/stop/api"
	"github.com/redmejia/stop/database"
)

func main() {
	// connection postgresql
	db, err := database.Connection()
	if err != nil {
		fmt.Println("ERROR", err)
	}

	// loggers
	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errors := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	success := log.New(os.Stdout, "SUCCESS\t", log.Ldate|log.Ltime|log.Lshortfile)

	server := log.New(os.Stdout, "SERVICE\t", log.Ldate|log.Ltime)

	api := api.Application{
		DB:      db,
		Info:    info,
		Error:   errors,
		Success: success,
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&api),
	}

	server.Println("Runing on http://127.0.0.1:8080")

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("err ", err)
	}
}
