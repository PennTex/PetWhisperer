package api

import (
	"log"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		Debug: true,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			secret := []byte(os.Getenv("AUTH0_CLIENT_SECRET"))

			if len(secret) == 0 {
				log.Fatal("AUTH0_CLIENT_SECRET is not set")
			}

			return secret, nil
		},
	})

	router.Handle("/pets", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(getPets)),
	))

	return router
}
