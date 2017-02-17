package app

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := struct{}{}

	RenderTemplate(w, "index", data)
}
