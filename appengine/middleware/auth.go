package middleware

import (
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var authKey = os.Getenv("AUTHORIZATION_KEY")

// Auth - Authentication middleware for communication between services
func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		if r.Header.Get("x-auth") != authKey {
			log.Infof(ctx, "%s != %s", r.Header.Get("x-auth"), authKey)

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Service Not Authorized"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
