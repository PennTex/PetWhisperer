// Logs you into Pet Whisperer.
// Useful for local development when needing an access token for API calls.
// Usage: login <username> <password>
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	auth0ClientID := "ys9ncsImAicfvkbsHbDnpF2aZ63uur3N"
	auth0ClientSecret := os.Getenv("PET_WHISPERER_AUTH0_CLI_CLIENT_SECRET")

	username := os.Args[1]
	password := os.Args[2]

	url := "https://pet-whisperer.auth0.com/oauth/token"

	payload := strings.NewReader("{\"grant_type\":\"password\",\"username\": \"" + username + "\",\"password\": \"" + password + "\",\"audience\": \"api.pet-whisperer.com\", \"client_id\": \"" + auth0ClientID + "\", \"client_secret\": \"" + auth0ClientSecret + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
