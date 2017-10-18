package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/appengine/middlewares"
	"github.com/PennTex/pet-whisperer/src/activityservice"
)

func init() {
	http.Handle("/", middleware.Auth(activityservice.New()))
}
