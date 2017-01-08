package main

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/repositories"
	"github.com/PennTex/PetWhisperer/services"
	"github.com/gorilla/mux"
)

var animalRepo repositories.MemoryRepo
var animalService = services.NewAnimalService(animalRepo)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/animals", getAnimals)
	r.HandleFunc("/animals/{animalID}", getAnimalByID)
	http.ListenAndServe(":8080", r)
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	animals := animalService.GetAnimals()
	message, err := json.Marshal(animals)

	if err != nil {
		panic(err)
	}

	w.Write(message)
}

func getAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
	animal := animalService.GetAnimalByID(animalID)
	message, err := json.Marshal(animal)

	if err != nil {
		panic(err)
	}

	w.Write(message)
}
