package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
	animals := animalService.GetAnimals()
	sendResponse(w, r, http.StatusOK, animals)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	animal := animalService.GetAnimal(animalID)
	sendResponse(w, r, http.StatusOK, animal)
}
