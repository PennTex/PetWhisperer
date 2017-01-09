package main

import (
	"encoding/json"
	"net/http"

	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/PennTex/PetWhisperer/repositories"
	"github.com/PennTex/PetWhisperer/services"
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
	r := mux.NewRouter()

	r.HandleFunc("/animals", getAnimals)
	r.HandleFunc("/animals/{animalID}", getAnimal)
	r.HandleFunc("/animals/{animalID}/activity", getAnimalActivity).
		Methods("GET")
	r.HandleFunc("/animals/{animalID}/activity", createActivity).
		Methods("POST")

	http.ListenAndServe(":8080", r)
}

func sendResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	response := Response{
		Data: data,
	}
	message, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	animals := animalService.GetAnimals()
	sendResponse(w, r, animals)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
	animal := animalService.GetAnimal(animalID)

	sendResponse(w, r, animal)
}

func getAnimalActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
	activities := activityService.GetAnimalActivity(animalID)

	sendResponse(w, r, activities)
}

func createActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animalID := vars["animalID"]
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
		theActivity := activity.NewFeedActivity(animalID, activityPost.FedBy)
		_ = activityService.CreateActivity(theActivity)
	} else if activityPost.Type == "medication" {
		theActivity := activity.NewMedicationActivity(animalID, activityPost.GivenBy)
		_ = activityService.CreateActivity(theActivity)
	}
}
