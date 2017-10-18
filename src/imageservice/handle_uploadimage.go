package imageservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
)

func handleUploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	f, fh, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		sendResponse(w, r, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		sendResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	// random filename, retaining existing extension.
	name := uuid.NewV4().String() + path.Ext(fh.Filename)

	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		sendResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		sendResponse(w, r, http.StatusInternalServerError, err)
		return
	}
	defer client.Close()

	log.Infof(ctx, "Using bucket name: %v\n\n", bucket)

	objectWriter := client.Bucket(bucket).Object(name).NewWriter(ctx)
	objectWriter.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	objectWriter.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	objectWriter.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(objectWriter, f); err != nil {
		log.Errorf(ctx, "failed to copy file into bucket: %v", err)
		sendResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if err := objectWriter.Close(); err != nil {
		log.Errorf(ctx, "failed closing object writer: %v", err)
		sendResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, name)
	log.Infof(ctx, "All %s", "success")

	sendResponse(w, r, http.StatusOK, publicURL)
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
