package v1

import "github.com/gorilla/mux"

func RegisterRoutes(router mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/v1").Subrouter()
  
  subRouter.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimalActivities).
    Methods("GET")
  subRouter.HandleFunc("/", postActivity).
    Methods("POST")

	return &router
}
