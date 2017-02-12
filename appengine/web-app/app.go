package app

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init() {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/callback", CallbackHandler)

	http.Handle("/", router)
}
