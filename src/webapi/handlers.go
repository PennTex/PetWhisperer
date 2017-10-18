package webapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"goengine"

	"github.com/gorilla/mux"
)

func handleGetPets(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(animalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/users/%s/animals", r.Context().Value("userID").(string))

	proxy.ServeHTTP(w, r)
}

func handlePostPet(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(animalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	animalReq := animalPostReq{}

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

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = "/animals"
	r.ContentLength = int64(buf.Len())
	r.Body = ioutil.NopCloser(buf)

	proxy.ServeHTTP(w, r)
}

func handleDeletePet(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(animalServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/animals/%s", animalID)

	proxy.ServeHTTP(w, r)
}

func handleGetPetsActivities(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(activityServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/%s", animalID)

	proxy.ServeHTTP(w, r)
}

func handlePostPetActivity(w http.ResponseWriter, r *http.Request) {
	animalID := mux.Vars(r)["animalID"]
	url, _ := url.Parse(activityServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	activityReq := activityPostReq{}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &activityReq)

	activityReq.By = r.Context().Value("userID").(string)

	activityAsJSON, err := json.Marshal(activityReq)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(activityAsJSON)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = fmt.Sprintf("/%s", animalID)
	r.ContentLength = int64(buf.Len())
	r.Body = ioutil.NopCloser(buf)

	proxy.ServeHTTP(w, r)
}

func handlePostImage(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(imageServiceBasePath)
	proxy := goengine.NewSingleHostReverseProxy(url)

	r.Header.Add("x-auth", servicesAuthorizationKey)
	r.URL.Path = "/upload"

	proxy.ServeHTTP(w, r)
}
