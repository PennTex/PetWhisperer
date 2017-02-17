package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animal/models"
	"github.com/PennTex/PetWhisperer/src/animal/repositories"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

var animalRepo repositories.CloudDatastoreRepository

type Response struct {
	Data interface{} `json:"data"`
}

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

// TODO: why must i pass a context :(
func getAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animals, _ := animalRepo.Get(ctx)
	sendResponse(w, r, http.StatusOK, animals)
}

func postAnimal(w http.ResponseWriter, r *http.Request) {
	var animal models.Animal
	b, _ := ioutil.ReadAll(r.Body)
	ctx := appengine.NewContext(r)

	json.Unmarshal(b, &animal)

	animalID, err := animalRepo.Create(ctx, animal)
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
	} else {
		sendResponse(w, r, http.StatusOK, animalID)
	}
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]
	animal, _ := animalRepo.GetByID(ctx, animalID)
	sendResponse(w, r, http.StatusOK, animal)
}
