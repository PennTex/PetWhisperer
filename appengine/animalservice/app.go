package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/appengine/middlewares"
	"github.com/PennTex/pet-whisperer/src/animalservice"
)

func init() {
	http.Handle("/", middleware.Auth(animalservice.New()))
}
