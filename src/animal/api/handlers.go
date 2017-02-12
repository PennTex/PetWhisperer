package api

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animal/repositories"
	"github.com/gorilla/mux"
)

var animalRepo repositories.InMemoryAnimalRepository

func sendResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	response := Response{
		Data: data,
	}
	message, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(message)
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	animals := animalRepo.GetAll()
	sendResponse(w, r, http.StatusOK, animals)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	animal := animalRepo.Get(animalID)
	sendResponse(w, r, http.StatusOK, animal)
}
