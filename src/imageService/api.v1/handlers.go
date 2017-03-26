package api

import (
	"encoding/json"
	"net/http"
)

type PetInfo struct {
	Type string `json:"type"`
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
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
