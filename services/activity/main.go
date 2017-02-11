package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PennTex/PetWhisperer/services/activity/repositories"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var activityRepo repositories.InMemoryActivityRepository
var activityService = NewActivitiesService(activityRepo)

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", getAnimalActivities).
		Methods("GET")
	r.HandleFunc("/", postActivity).
		Methods("POST")

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.RecoveryHandler()(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
