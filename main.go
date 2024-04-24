package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sgrkabadi96/bookAppClone/routes"
)

func main() {
	r := mux.NewRouter()
	routes.HandleUserRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":4000", r)
}
