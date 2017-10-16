package middleware

import (
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Auth - Authentication middleware for communication between services
func Auth(h http.Handler) http.Handler {
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
