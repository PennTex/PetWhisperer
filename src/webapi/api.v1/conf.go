package api

import "os"

var (
	AnimalServiceBasePath    string
	Auth0ClientSecret        string
	ServicesAuthorizationKey string
)

func init() {
	AnimalServiceBasePath = os.Getenv("ANIMAL_SERVICE_BASE_PATH")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	ServicesAuthorizationKey = os.Getenv("AUTHORIZATION_KEY")
}
