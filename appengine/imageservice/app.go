package app

import (
	"net/http"

	"github.com/PennTex/PetWhisperer/src/imageservice/api.v1"
)

func init() {
	http.Handle("/v1/", api.New())
}
