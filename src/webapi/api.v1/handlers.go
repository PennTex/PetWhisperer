package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PennTex/pet-whisperer/src/webapi/goengine"
	"github.com/gorilla/mux"
)

func getPets(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(AnimalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/users/%s/animals", r.Context().Value("userID").(string))

	proxy.ServeHTTP(w, r)
}

func postPet(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(AnimalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	animalReq := AnimalPostReq{}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalReq.Owners = []string{
		r.Context().Value("userID").(string),
	}

	animalAsJSON, err := json.Marshal(animalReq)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(animalAsJSON)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = "/animals"
	r.ContentLength = int64(buf.Len())
	r.Body = ioutil.NopCloser(buf)

	proxy.ServeHTTP(w, r)
}

func deletePet(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(AnimalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/animals/%s", animalID)

	proxy.ServeHTTP(w, r)
}

func getPetsActivities(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(ActivityServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/%s", animalID)

	proxy.ServeHTTP(w, r)
}

func postPetActivity(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(ActivityServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	activityReq := ActivityPostReq{}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &activityReq)

	activityReq.By = r.Context().Value("userID").(string)

	activityAsJSON, err := json.Marshal(activityReq)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(activityAsJSON)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/%s", animalID)
	r.ContentLength = int64(buf.Len())
	r.Body = ioutil.NopCloser(buf)

	proxy.ServeHTTP(w, r)
}

func postImage(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(ImageServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", ServicesAuthorizationKey)
	r.URL.Path = "/upload"

	proxy.ServeHTTP(w, r)
}
