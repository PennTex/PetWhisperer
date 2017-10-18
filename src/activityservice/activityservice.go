package activityservice

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", handleGetAnimalActivities).Methods("GET")
	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", handlePostActivity).Methods("POST")

	return router
}
