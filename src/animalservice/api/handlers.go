package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine/log"

	"github.com/PennTex/pet-whisperer/src/animalservice/models"
	"github.com/PennTex/pet-whisperer/src/animalservice/repositories"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var animalRepo repositories.CloudDatastoreRepository

func getUsersAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ownerID := mux.Vars(r)["userID"]

	animals, err := animalRepo.GetByOwnerID(ctx, ownerID)

	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(ctx, "Animals retrieved: %s", animals)

	sendResponse(w, r, http.StatusOK, animals)
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalIDs := r.URL.Query()["animalID"]

	if animalIDs != nil {
		var animals []models.Animal

		for _, animalID := range animalIDs {
			animal, err := animalRepo.GetByID(ctx, animalID)

			if err != nil {
				sendResponse(w, r, http.StatusInternalServerError, err.Error())
				return
			}

			animals = append(animals, *animal)
		}

		sendResponse(w, r, http.StatusOK, animals)
	} else {
		animals, err := animalRepo.Get(ctx)

		if err != nil {
			sendResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		sendResponse(w, r, http.StatusOK, animals)
	}
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]

	animal, err := animalRepo.GetByID(ctx, animalID)

	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, r, http.StatusOK, animal)
}

func postAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalReq := AnimalPostReq{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalID, err := animalRepo.Create(ctx, &models.Animal{
		Typ:       animalReq.Typ,
		Name:      animalReq.Name,
		Birthday:  animalReq.Birthday,
		Owners:    animalReq.Owners,
		ImageURL:  animalReq.ImageURL,
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof(ctx, "New animal ID: %s", animalID)
	sendResponse(w, r, http.StatusOK, animalID)
}

func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]

	err := animalRepo.Destroy(ctx, animalID)
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, r, http.StatusNoContent, nil)
}

func sendResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	message, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	w.Write(message)
}
