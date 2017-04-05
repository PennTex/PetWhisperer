package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	n := negroni.New(
		negroni.HandlerFunc(corsMiddleware),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.HandlerFunc(userMiddleware),
	)

	router := mux.NewRouter()
	router.HandleFunc("/pets", getPets).Methods("GET")
	router.HandleFunc("/pets", postPet).Methods("POST")
	router.HandleFunc("/pets/{animalID:[0-9a-z-]{36}}", deletePet).Methods("DELETE")

	router.HandleFunc("/images", postImage).Methods("POST")

	n.UseHandler(router)

	return n
}
