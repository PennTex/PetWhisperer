package app

import (
	"net/http"

	api "github.com/PennTex/pet-whisperer/src/webapi/api"
)

func init() {
	http.Handle("/", api.New())
}
