// Package webapi is the client facing api that interacts with all of the backend services
package webapi

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// New returns the webapi handler
func New() http.Handler {
	n := negroni.New(
		negroni.HandlerFunc(corsMiddleware),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.HandlerFunc(userMiddleware),
	)

	router := mux.NewRouter()
	router.HandleFunc("/pets", handleGetPets).Methods("GET")
	router.HandleFunc("/pets", handlePostPet).Methods("POST")
	router.HandleFunc("/pets/{animalID:[0-9a-z-]{36}}", handleDeletePet).Methods("DELETE")
	router.HandleFunc("/pets/{animalID:[0-9a-z-]{36}}/activities", handleGetPetsActivities).Methods("GET")
	router.HandleFunc("/pets/{animalID:[0-9a-z-]{36}}/activities", handlePostPetActivity).Methods("POST")

	router.HandleFunc("/images", handlePostImage).Methods("POST")

	n.UseHandler(router)

	return n
}
