package main

import (
	"net/http"

	"github.com/redmejia/stop/api"
)

func routes(a *api.Application) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/pants", a.ProductPants)
	mux.HandleFunc("/shirts", a.ProductShirt)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", http.StripPrefix("/", fs))

	return mux
}
