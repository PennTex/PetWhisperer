package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/src/imageservice/api.v1"
)

func init() {
	http.Handle("/", api.New())
}
