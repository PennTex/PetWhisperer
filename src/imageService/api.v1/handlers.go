package api

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type PetInfo struct {
	Type string `json:"type"`
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("image")
	if err != nil {
		log.Errorf(ctx, err.Error(), nil)
		return
	}

	defer file.Close()

	sendResponse(w, r, http.StatusOK, struct {
		PetInfo  PetInfo `json:"petInfo"`
		ImageURL string  `json:"image_url"`
	}{
		PetInfo: PetInfo{
			Type: "dog",
		},
		ImageURL: "https://images-na.ssl-images-amazon.com/images/G/01/img15/pet-products/small-tiles/23695_pets_vertical_store_dogs_small_tile_8._CB312176604_.jpg",
	})
}

func sendResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	message, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	w.Write(message)
}
