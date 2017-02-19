package app

import (
	"encoding/gob"
	"net/http"

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
}
