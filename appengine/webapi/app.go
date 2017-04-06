package app

import (
	"net/http"

	api "github.com/PennTex/pet-whisperer/src/webapi/api.v1"
)

func init() {
	http.Handle("/", api.New())
}
