package app

import (
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animalservice/api"
)

func init() {
	router := api.GetRouter()
	http.Handle("/", router)
}
