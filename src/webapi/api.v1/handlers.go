package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/context"

	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func getPets(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/animals", AnimalServiceBasePath, context.Get(r, "userID")), nil)
	req.Header.Add("Authentication", os.Getenv("AUTHORIZATION_KEY"))

	response, err := client.Do(req)
	if err != nil {
		log.Criticalf(ctx, "could not get animals from animals service: %s", err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Criticalf(ctx, "could not read response from animals service: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write(responseData)
}

func postPet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	animalReq := AnimalPostReq{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &animalReq)

	animalReq.Owners = []string{
		context.Get(r, "userID").(string),
	}

	animalAsJSON, err := json.Marshal(animalReq)
	if err != nil {
		panic(err)
	}

	log.Debugf(ctx, "webapi: calling %s", fmt.Sprintf("%s/animals", AnimalServiceBasePath))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/animals", AnimalServiceBasePath), bytes.NewReader(animalAsJSON))
	req.Header.Add("Authorization", os.Getenv("AUTHORIZATION_KEY"))

	response, err := client.Do(req)
	if err != nil {
		log.Criticalf(ctx, "could not post animal to animals service: %s", err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Criticalf(ctx, "could not read response from animals service: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write(responseData)
}
