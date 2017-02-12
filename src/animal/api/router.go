package api

import (
	"github.com/PennTex/PetWhisperer/src/animal/api/v1"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routerWithV1Routes := v1.RegisterRoutes(*router)
	return routerWithV1Routes
}
