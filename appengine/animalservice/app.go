package app

import (
	"net/http"

	"os"

	"github.com/PennTex/PetWhisperer/src/animalservice"
	"github.com/PennTex/PetWhisperer/src/animalservice/api.v1"
	"github.com/PennTex/PetWhisperer/src/animalservice/repositories"
	"github.com/gorilla/mux"
)

func init() {
	animalRepository := repositories.CloudDatastoreRepository{}
	animalService := animalservice.NewAnimalService(animalRepository)
	animalAPI := api.NewAnimalAPI(animalService)

	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/v1").Subrouter()

	subRouter.HandleFunc("/", animalAPI.PostAnimal).
		Methods("POST")

	subRouter.HandleFunc("/", animalAPI.GetAnimals).
		Methods("GET")
	subRouter.HandleFunc("/{animalID:[0-9a-z-]{36}}", animalAPI.GetAnimal).
		Methods("GET")
	subRouter.HandleFunc("/user/{userID}/animals", animalAPI.GetUsersAnimals).
		Methods("GET")

	http.Handle("/", AuthMiddleware(router))
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != os.Getenv("AUTHORIZATION_KEY") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
