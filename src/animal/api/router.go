package api

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", getAnimals)
	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimal)

	return router
}
