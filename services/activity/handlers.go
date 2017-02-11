package main

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/services/activity/models"
	"github.com/PennTex/PetWhisperer/services/activity/repositories"
	"github.com/gorilla/mux"
)

var activityRepo repositories.InMemoryActivityRepository

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

func getAnimalActivities(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	activities := activityRepo.GetByAnimalID(animalID)

	sendResponse(w, r, http.StatusOK, activities)
}

func postActivity(w http.ResponseWriter, r *http.Request) {
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
		err = activityRepo.Create(models.NewFeedActivity(animalID, activityPost.FedBy))
	} else if activityPost.Type == "medication" {
		err = activityRepo.Create(models.NewMedicationActivity(animalID, activityPost.GivenBy))
	} else {
		sendResponse(w, r, http.StatusBadRequest, "Activity type not supported.")
		return
	}

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}
