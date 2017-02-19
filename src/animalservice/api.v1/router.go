package api

import (
	"github.com/PennTex/PetWhisperer/src/animalservice"
	"github.com/PennTex/PetWhisperer/src/animalservice/repositories"
	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router) *mux.Router {
	animalRepository := repositories.CloudDatastoreRepository{}
	animalService := animalservice.NewAnimalService(animalRepository)
	animalAPI := NewAnimalAPI(animalService)

	router.HandleFunc("/", animalAPI.PostAnimal).
		Methods("POST")

	router.HandleFunc("/", animalAPI.GetAnimals).
		Methods("GET")
	router.HandleFunc("/{animalID:[0-9a-z-]{36}}", animalAPI.GetAnimal).
		Methods("GET")
	router.HandleFunc("/users/{userID}/animals", animalAPI.GetUsersAnimals).
		Methods("GET")

	return router
}
