package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PennTex/PetWhisperer/src/webapi/goengine"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func getPets(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	url, _ := url.Parse(AnimalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	log.Infof(ctx, "Services Auth Key: %s", ServicesAuthorizationKey)

	r.Header.Add("Authorization", ServicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/users/%s/animals", context.Get(r, "userID"))

	proxy.ServeHTTP(w, r)
}

func postPet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	url, _ := url.Parse(AnimalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	log.Infof(ctx, "Services Auth Key: %s", ServicesAuthorizationKey)

	animalReq := AnimalPostReq{}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalReq.Owners = []string{
		context.Get(r, "userID").(string),
	}

	animalAsJSON, err := json.Marshal(animalReq)
	if err != nil {
		panic(err)
	}

	log.Infof(ctx, "animal bytes %s", animalAsJSON)

	buf := bytes.NewBuffer(animalAsJSON)

	r.Header.Add("Authorization", ServicesAuthorizationKey)
	r.URL.Path = "/animals"
	r.ContentLength = int64(buf.Len())
	r.Body = ioutil.NopCloser(buf)

	proxy.ServeHTTP(w, r)
}

func deletePet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	animalID := mux.Vars(r)["animalID"]

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/animals/%s", AnimalServiceBasePath, animalID), nil)
	req.Header.Add("Authentication", ServicesAuthorizationKey)

	response, err := client.Do(req)
	if err != nil {
		log.Criticalf(ctx, "could not delete animal using AnimalService: %s", err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Criticalf(ctx, "could not read response from AnimalService: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write(responseData)
}

func postImage(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(ImageServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("Authorization", ServicesAuthorizationKey)
	r.URL.Path = "/upload"

	proxy.ServeHTTP(w, r)
}
