package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/appengine/middleware"
	"github.com/PennTex/pet-whisperer/src/animalservice/api.v1"
)

func init() {
	http.Handle("/", middleware.Auth(api.New()))
}
