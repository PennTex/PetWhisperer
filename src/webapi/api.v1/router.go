package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// create common middleware to be shared across routes
	common := negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.HandlerFunc(userMiddleware),
	)

	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.WriteHeader(http.StatusNoContent)

			w.Write([]byte(""))
		})

	router.Handle("/pets", common.With(
		negroni.Wrap(http.HandlerFunc(getPets)),
	)).Methods("GET")

	router.Handle("/pets", common.With(
		negroni.Wrap(http.HandlerFunc(postPet)),
	)).Methods("POST")

	return router
}
