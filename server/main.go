package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PennTex/PetWhisperer/repositories"
	"github.com/PennTex/PetWhisperer/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var animalRepo repositories.InMemoryAnimalRepo
var animalService = services.NewAnimalService(animalRepo)
var activityRepo repositories.InMemoryActivityRepo
var activityService = services.NewActivityService(activityRepo)

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/animals", getAnimals)
	r.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}", getAnimal)
	r.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}/activity", getAnimalActivity).
		Methods("GET")
	r.HandleFunc("/animals/{animalID:[0-9a-z-]{36}}/activity", createActivity).
		Methods("POST")

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.RecoveryHandler()(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
