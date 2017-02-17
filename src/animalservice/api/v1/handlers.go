package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	"cloud.google.com/go/storage"

	"github.com/PennTex/PetWhisperer/src/animalservice"
	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	"github.com/PennTex/PetWhisperer/src/animalservice/repositories"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine"
)

var animalRepo repositories.CloudDatastoreRepository

type Response struct {
	Data interface{} `json:"data"`
}

func sendResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	response := Response{
		Data: data,
	}
	message, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(message)
}

// TODO: why must i pass a context :(
func getAnimals(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animals, _ := animalRepo.Get(ctx)
	sendResponse(w, r, http.StatusOK, animals)
}

func postAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	var animal models.Animal
	animal.CreatedAt = time.Now().Unix()
	animal.ImageURL, _ = uploadFileFromForm(r)

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animal)

	animalID, err := animalRepo.Create(ctx, animal)
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err.Error())
	} else {
		sendResponse(w, r, http.StatusOK, animalID)
	}
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	animalID := mux.Vars(r)["animalID"]
	animal, _ := animalRepo.GetByID(ctx, animalID)
	sendResponse(w, r, http.StatusOK, animal)
}

func uploadFileFromForm(r *http.Request) (url string, err error) {
	ctx := appengine.NewContext(r)
	client, err := storage.NewClient(ctx)

	bucket := client.Bucket(animalservice.StorageBucketName)

	f, fh, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	// random filename, retaining existing extension.
	name := uuid.NewV4().String() + path.Ext(fh.Filename)

	w := bucket.Object(name).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, f); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	return fmt.Sprintf(publicURL, animalservice.StorageBucketName, name), nil
}
