package v1

import "github.com/gorilla/mux"

func RegisterRoutes(router mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/", getAnimals).Methods("GET")
	subRouter.HandleFunc("/", postAnimal).Methods("POST")
	subRouter.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimal).Methods("GET")
	return &router
}
