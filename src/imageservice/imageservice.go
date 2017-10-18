package imageservice

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/upload", handleUploadImage).Methods("POST")

	return router
}
