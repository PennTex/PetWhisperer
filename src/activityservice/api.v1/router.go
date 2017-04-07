package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimalActivities).Methods("GET")
	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", postActivity).Methods("POST")

	return router
}
