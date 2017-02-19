package app

import (
	"net/http"

	"os"

	"github.com/PennTex/PetWhisperer/src/animalservice/api.v1"
	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter().StrictSlash(true)
	_ = api.NewRouter(router)

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
