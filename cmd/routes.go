package main

import (
	"net/http"

	"github.com/redmejia/stop/api"
	"github.com/redmejia/stop/cors"
)

func routes(a *api.Application) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/pants", a.ProductPants)
	mux.HandleFunc("/shirts", a.ProductShirt)
	mux.HandleFunc("/arrivals", a.ProductArrivals)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", http.StripPrefix("/", fs))

	return cors.Cors(mux)
}
