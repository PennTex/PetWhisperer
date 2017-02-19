package app

import (
	"net/http"

	"os"

	"github.com/PennTex/PetWhisperer/src/animalservice/api.v1"
	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter().StrictSlash(true)

	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router = api.NewRouter(v1Router)

	http.Handle("/", authMiddleware(router))
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != os.Getenv("AUTHORIZATION_KEY") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
