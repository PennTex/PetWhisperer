package app

import (
	"net/http"

	"google.golang.org/appengine"

	"os"

	"github.com/PennTex/pet-whisperer/src/activityservice/api.v1"
	"google.golang.org/appengine/log"
)

func init() {
	http.Handle("/", authMiddleware(api.New()))
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		if r.Header.Get("x-auth") != os.Getenv("AUTHORIZATION_KEY") {
			log.Infof(ctx, "%s != %s", r.Header.Get("x-auth"), os.Getenv("AUTHORIZATION_KEY"))

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Service Not Authorized"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
