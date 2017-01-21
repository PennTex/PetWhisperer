package main

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/models/activity"
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

func getAnimalActivity(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	activities := activityService.GetAnimalActivity(animalID)

	sendResponse(w, r, http.StatusOK, activities)
}

func createActivity(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	decoder := json.NewDecoder(r.Body)

	var activityPost struct {
		Type    string `json:"type"`
		FedBy   string `json:"fed_by"`
		GivenBy string `json:"given_by"`
	}
	err := decoder.Decode(&activityPost)

	if err != nil {
		panic(err)
	}

	if activityPost.Type == "feed" {
		err = activityService.CreateActivity(activity.NewFeedActivity(animalID, activityPost.FedBy))
	} else if activityPost.Type == "medication" {
		err = activityService.CreateActivity(activity.NewMedicationActivity(animalID, activityPost.GivenBy))
	} else {
		sendResponse(w, r, http.StatusBadRequest, "Activity type not supported.")
		return
	}

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}
