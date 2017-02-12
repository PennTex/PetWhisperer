package activity

import (
	"net/http"

	"github.com/PennTex/PetWhisperer/src/activity/api"
)

func init() {
	router := api.NewRouter()

	http.Handle("/", router)
}
