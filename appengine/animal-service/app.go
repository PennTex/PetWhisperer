package app

import (
	"net/http"

	"github.com/PennTex/PetWhisperer/src/animal/api"
)

func init() {
	router := api.NewRouter()
	http.Handle("/", router)
}
