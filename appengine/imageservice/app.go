package app

import (
	"net/http"

	"github.com/PennTex/PetWhisperer/src/imageservice/api.v1"
)

func init() {
	http.Handle("/", api.New())
}
