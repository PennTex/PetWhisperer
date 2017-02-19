package app

import (
	"net/http"

	api "github.com/PennTex/PetWhisperer/src/webapi/api.v1"
)

func init() {
	router := api.NewRouter()
	http.Handle("/", router)
}
