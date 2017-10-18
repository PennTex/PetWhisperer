package activityservice

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"

	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

var activityRepo cloudDatastoreRepository

func sendResponse(w http.ResponseWriter, r *http.Request, status int, errorPayload interface{}, dataPayload interface{}) {
	message, err := json.Marshal(response{
		Error: errorPayload,
		Data:  dataPayload,
	})

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(message)
}

func handleGetAnimalActivities(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	animalID := mux.Vars(r)["animalID"]
	activities, _ := activityRepo.getByAnimalID(ctx, animalID)

	log.Infof(ctx, "Activities retrieved : %v", activities)

	sendResponse(w, r, http.StatusOK, nil, activities)
}

func handlePostActivity(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]
	decoder := json.NewDecoder(r.Body)

	var activityPost struct {
		Type string `json:"type"`
		By   string `json:"by"`
		At   int64  `json:"at"`
		Note string `json:"note"`
	}

	err := decoder.Decode(&activityPost)

	if err != nil {
		sendResponse(w, r, http.StatusBadRequest, "Invalid activity payload supplied.", nil)
		return
	}

	if activityPost.Type != "feed" && activityPost.Type != "medication" {
		sendResponse(w, r, http.StatusBadRequest, "Activity type not supported.", nil)
		return
	}

	createdActivity, err := activityRepo.create(ctx, &activity{
		Typ:       activityPost.Type,
		AnimalID:  animalID,
		By:        activityPost.By,
		Note:      activityPost.Note,
		At:        activityPost.At,
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		panic(err)
	}

	sendResponse(w, r, http.StatusCreated, nil, createdActivity)
}
