package webapi

import (
	"context"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware

const publicKey = `-----BEGIN CERTIFICATE-----
MIIC9jCCAd6gAwIBAgIJKkDMgJzuetGBMA0GCSqGSIb3DQEBBQUAMCIxIDAeBgNV
BAMTF3BldC13aGlzcGVyZXIuYXV0aDAuY29tMB4XDTE3MDIxMjE2NTYxNVoXDTMw
MTAyMjE2NTYxNVowIjEgMB4GA1UEAxMXcGV0LXdoaXNwZXJlci5hdXRoMC5jb20w
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDXdS/tCe55CU23mmqQ0A0k
gAyKf/fWgbdSHIPCxBad+AUMTRKpRa5VAJwpkywB5JODLpuZDHtz3dT5eNPRZj8N
z/okqa/d8eesTrVjcVaBV8S+4mcT9QkInyN80eLjHuO6EfplrJzNUPvfKvVE7HiR
OHF+qxK40Rd34LWoG0OpB+ZZMsW9UGYQod6RR1Sjkh7ZzdcImymJe/qEz+n5I0fo
h6vLJBtqpy80y2vEmbLjO4s9OufYpcBm0GlJRXoavINv1rTRyA3MnU6zWm3ZFNpD
WAnEkLULCNOMeSLYDFd9WAC0VLXeM/hk7GAjDFEN14ITDLs/jTQT0RNh5W0r0AfH
AgMBAAGjLzAtMAwGA1UdEwQFMAMBAf8wHQYDVR0OBBYEFCTRd2ettLOtZ/1qdXk4
7trkyeKHMA0GCSqGSIb3DQEBBQUAA4IBAQDJESqeZ5yagUdrak2UuMmrVfh1JO54
lUPidj4KqNM+YEdywBvw4Jwg9Zj5hqnNdvzCqUYyWoKrFfAOUXu3LX3VKYsXvwbl
eSdXiSNdeSff2oooV9BJi8oI2kBIHHU+ip5wLxNZEit4W5h8kr4j0umbXenPR0X1
UuH0wMJqVstPKEkXbzOm32sN2L63kCHQDETZGnXeRcbz0i7zZBjU8L5BkRSUIJ++
AtKgZIfdi9sH01sOhZs5bam54wBQl2WxunDIKLy7f0beHFRVPrOeqOto1I1dJNe1
01QtdTA/DZ+m5copG0N99+dE2LfH0ULIlJNLXjUMVHHnSsDomuu7GAYI
-----END CERTIFICATE-----`

func init() {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		Debug: true,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		},
		SigningMethod: jwt.SigningMethodRS256,
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
