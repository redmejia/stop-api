package api

import (
	"fmt"
	"net/http"
)

func (a *Application) ProductPants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func (a *Application) ProductShirt(w http.ResponseWriter, r *http.Request) {
	// not yet
}
