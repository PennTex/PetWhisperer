package animalservice

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/animals", postAnimal).Methods("POST")
	router.HandleFunc("/animals", getAnimals).Methods("GET")
	router.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}", getAnimal).Methods("GET")
	router.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}", deleteAnimal).Methods("DELETE")
	router.HandleFunc("/users/{userID}/animals", getUsersAnimals).Methods("GET")

	return router
}
