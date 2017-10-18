package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/src/webapi"
)

func init() {
	http.Handle("/", webapi.New())
}
