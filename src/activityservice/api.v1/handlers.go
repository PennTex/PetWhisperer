package api

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"

	"github.com/PennTex/pet-whisperer/src/activityservice/models"
	"github.com/PennTex/pet-whisperer/src/activityservice/repositories"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

var activityRepo repositories.CloudDatastoreRepository

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
	ctx := appengine.NewContext(r)

	animalID := mux.Vars(r)["animalID"]
	activities, _ := activityRepo.GetByAnimalID(ctx, animalID)

	log.Infof(ctx, "Activities retrieved : %v", activities)

	sendResponse(w, r, http.StatusOK, activities)
}

func postActivity(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]
	decoder := json.NewDecoder(r.Body)

	var activityPost struct {
		Type string `json:"type"`
		By   string `json:"by"`
		At   string `json:"at"`
		Note string `json:"note"`
	}
	err := decoder.Decode(&activityPost)

	if err != nil {
		panic(err)
	}

	var createdActivity *models.Activity

	if activityPost.Type == "feed" {
		activity := models.NewActivity("feed", animalID, activityPost.By)
		activity.Note = activityPost.Note
		createdActivity, err = activityRepo.Create(ctx, activity)
	} else if activityPost.Type == "medication" {
		activity := models.NewActivity("medication", animalID, activityPost.By)
		activity.Note = activityPost.Note
		createdActivity, err = activityRepo.Create(ctx, activity)
	} else {
		sendResponse(w, r, http.StatusBadRequest, "Activity type not supported.")
		return
	}

	if err != nil {
		panic(err)
	}

	sendResponse(w, r, http.StatusCreated, createdActivity)
}
