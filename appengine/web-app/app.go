package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler)
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.Handle("/", router)
}
