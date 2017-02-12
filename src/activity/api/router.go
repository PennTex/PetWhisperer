package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimalActivities).
		Methods("GET")
	router.HandleFunc("/", postActivity).
		Methods("POST")

	return router
}
