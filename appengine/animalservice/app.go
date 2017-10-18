package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/appengine/middlewares"
	"github.com/PennTex/pet-whisperer/src/animalservice/api"
)

func init() {
	http.Handle("/", middleware.Auth(api.New()))
}
