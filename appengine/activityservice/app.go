package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/src/activityservice/api.v1"
	"github.com/PennTex/pet-whisperer/src/middleware"
)

func init() {
	http.Handle("/", middleware.Auth(api.New()))
}
