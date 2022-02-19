package main

import (
	"log"
	"net/http"
	"os"

	"github.com/redmejia/stop/api"
)

func main() {

	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	server := log.New(os.Stdout, "SERVICE\t", log.Ldate|log.Ltime)

	api := api.Application{Info: info}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&api),
	}

	server.Println("Runing on http://127.0.0.1:8080")

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("err ", err)
	}
}
