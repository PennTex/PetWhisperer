package api

import (
	"github.com/PennTex/PetWhisperer/src/activity/api/v1"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routerWithV1Routes := v1.RegisterRoutes(*router)
	return routerWithV1Routes
}
