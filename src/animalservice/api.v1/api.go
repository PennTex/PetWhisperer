package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animalservice"
	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type AnimalAPI struct {
	AnimalService *animalservice.AnimalService
}

// NewAnimalAPI Creates a new Animal API
func NewAnimalAPI(service *animalservice.AnimalService) *AnimalAPI {
	return &AnimalAPI{
		AnimalService: service,
	}
}

// GetUsersAnimals Gets all animals owned by the given user
func (a *AnimalAPI) GetUsersAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ownerID := mux.Vars(r)["userID"]

	animals, err := a.AnimalService.GetAnimalsByOwnerID(ctx, ownerID)
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, r, http.StatusOK, animals)
}

// GetAnimals Gets all animals, or specific animals if animalID query param is passed in.
func (a *AnimalAPI) GetAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalIDs := r.URL.Query()["animalID"]

	if animalIDs != nil {
		var animals []models.Animal

		for _, animalID := range animalIDs {
			animal, err := a.AnimalService.GetAnimal(ctx, animalID)
			if err != nil {
				sendResponse(w, r, http.StatusInternalServerError, err.Error())
				return
			}

			animals = append(animals, *animal)
		}

		sendResponse(w, r, http.StatusOK, animals)
	} else {
		animals, err := a.AnimalService.GetAnimals(ctx)
		if err != nil {
			sendResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		sendResponse(w, r, http.StatusOK, animals)
	}
}

// GetAnimal Gets a single animal by its ID.
func (a *AnimalAPI) GetAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]

	animal, err := a.AnimalService.GetAnimal(ctx, animalID)
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, r, http.StatusOK, animal)
}

// PostAnimal Creates a new Animal
func (a *AnimalAPI) PostAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalReq := AnimalPostReq{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalID, err := a.AnimalService.CreateAnimal(ctx, &models.Animal{
		Typ:      animalReq.Typ,
		Name:     animalReq.Name,
		Birthday: animalReq.Birthday,
		Owners:   animalReq.Owners,
		ImageURL: animalReq.ImageURL,
	})

	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("New animal ID: %s", animalID)
	sendResponse(w, r, http.StatusOK, animalID)
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
