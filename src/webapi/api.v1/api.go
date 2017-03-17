package api

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	n := negroni.New(
		negroni.HandlerFunc(corsHandler),
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.HandlerFunc(userMiddleware),
	)

	router := mux.NewRouter()
	router.HandleFunc("/pets", getPets).Methods("GET")
	router.HandleFunc("/pets", postPet).Methods("POST")

	n.UseHandler(router)

	return n
}

func corsHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

	if r.Method == "OPTIONS" {
		return
	}

	next(w, r)
}
