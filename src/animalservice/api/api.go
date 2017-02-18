package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animalservice"
	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	"github.com/PennTex/PetWhisperer/src/animalservice/repositories"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type animalApi struct {
	animalService *animalservice.AnimalService
}

func RegisterHandlers(subRouter *mux.Router) *mux.Router {
	animalRepository := repositories.CloudDatastoreRepository{}
	animalService := animalservice.NewAnimalService(animalRepository)
	animalApi := animalApi{
		animalService: animalService,
	}

	subRouter.HandleFunc("/", animalApi.getAnimals)
	subRouter.HandleFunc("/{animalID:[0-9a-z-]{36}}", animalApi.getAnimal)

	return subRouter
}

func sendResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	response := models.Response{
		Error: nil,
		Data:  data,
	}

	message, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(message)
}

func (a *animalApi) getAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animals, _ := a.animalService.GetAnimals(ctx)
	sendResponse(w, r, http.StatusOK, animals)
}

func (a *animalApi) getAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]

	animal, _ := a.animalService.GetAnimal(ctx, animalID)
	sendResponse(w, r, http.StatusOK, animal)
}

func (a *animalApi) postAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	var animalReq struct {
		Typ      string   `json:"type"`
		Name     string   `json:"name"`
		Birthday int64    `json:"birthday"`
		Owners   []string `json:"owners"`
		ImageURL string   `json:"image_url"`
	}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalID, _ := a.animalService.CreateAnimal(ctx, &models.Animal{
		Typ:      animalReq.Typ,
		Name:     animalReq.Name,
		Birthday: animalReq.Birthday,
		Owners:   animalReq.Owners,
		ImageURL: animalReq.ImageURL,
	})

	sendResponse(w, r, http.StatusOK, animalID)
}
