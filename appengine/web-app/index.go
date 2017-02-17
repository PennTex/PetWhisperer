package app

import (
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		RedirectUrl string
	}{
		RedirectUrl: os.Getenv("AUTH0_CALLBACK_URL"),
	}

	RenderTemplate(w, "index", data)
}
