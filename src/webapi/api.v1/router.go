package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pets", getPets).
		Methods("GET")

	return router
}
