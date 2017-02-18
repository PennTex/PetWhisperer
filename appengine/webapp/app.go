package app

import (
	"encoding/gob"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

func init() {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/callback", CallbackHandler)
	http.Handle("/dashboard", negroni.New(
		negroni.HandlerFunc(IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(DashboardHandler)),
	))
}
