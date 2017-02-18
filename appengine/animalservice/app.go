package app

import (
	"net/http"

	animalserviceapi "github.com/PennTex/PetWhisperer/src/animalservice/api"
	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/v1").Subrouter()
	_ = animalserviceapi.RegisterHandlers(subRouter)
	http.Handle("/", router)
}
