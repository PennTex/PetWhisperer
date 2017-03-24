package app

import (
	"net/http"

	api "github.com/PennTex/PetWhisperer/src/webapi/api.v1"
)

func init() {
	http.Handle("/", api.New())
}
