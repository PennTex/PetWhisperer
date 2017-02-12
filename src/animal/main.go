package main

import (
	"log"
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animal/api"
)

func main() {
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
