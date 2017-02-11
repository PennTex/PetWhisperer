package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PennTex/PetWhisperer/services/animal/repositories"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var animalRepo repositories.InMemoryAnimalRepository
var animalService = NewAnimalService(animalRepo)

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/animals", getAnimals)
	r.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}", getAnimal)

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.RecoveryHandler()(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
