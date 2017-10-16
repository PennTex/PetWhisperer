package api

import (
	"context"
	"errors"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware

func init() {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		Debug: true,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			secret := []byte(Auth0ClientSecret)

			if len(secret) == 0 {
				return nil, errors.New("AUTH0_CLIENT_SECRET is not set")
			}

			return secret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func userMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := appengine.NewContext(r)
	user := r.Context().Value("user").(*jwt.Token)

	for k, v := range user.Claims.(jwt.MapClaims) {
		log.Infof(ctx, "Auth0 Claim %s: \t%#v\n", k, v)
	}

	userID := user.Claims.(jwt.MapClaims)["sub"]

	log.Infof(ctx, "UserID: %s", userID)

	newRequest := r.WithContext(context.WithValue(r.Context(), "userID", userID))
	*r = *newRequest

	next(w, r)
}

func corsMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")

	if r.Method == "OPTIONS" {
		return
	}

	next(w, r)
}
