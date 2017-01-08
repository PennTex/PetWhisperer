package main

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/repositories"
	"github.com/PennTex/PetWhisperer/services"
	"github.com/gorilla/mux"
)

var animalRepo repositories.InMemoryAnimalRepo
var animalService = services.NewAnimalService(animalRepo)
var activityRepo repositories.InMemoryActivityRepo
var activityService = services.NewActivityService(activityRepo)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/animals", getAnimals)
	r.HandleFunc("/animals/{animalID}", getAnimalByID)
	r.HandleFunc("/animals/{animalID}/activity", getAnimalActivity)

	http.ListenAndServe(":8080", r)
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	animals := animalService.GetAnimals()
	message, err := json.Marshal(animals)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func getAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
	animal := animalService.GetAnimal(animalID)
	message, err := json.Marshal(animal)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func getAnimalActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
	activities := activityService.GetAnimalActivity(animalID)
	message, err := json.Marshal(activities)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}
