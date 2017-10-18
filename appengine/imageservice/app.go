package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/src/imageservice/api"
)

func init() {
	http.Handle("/", api.New())
}
