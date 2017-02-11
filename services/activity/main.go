package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimalActivities).
		Methods("GET")
	r.HandleFunc("/", postActivity).
		Methods("POST")

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.RecoveryHandler()(router)

	fmt.Println("activity service listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
