package main

import (
	"net/http"

	"github.com/redmejia/stop/api"
)

func routes(a *api.Application) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", a.ProductPants)
	return mux
}
