package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Response struct {
	Data interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", getAnimals)
	r.HandleFunc("/{animalID:[0-9a-z-]{36}}", getAnimal)

	router := handlers.LoggingHandler(os.Stdout, r)
	router = handlers.RecoveryHandler()(router)

	fmt.Println("animal service listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
