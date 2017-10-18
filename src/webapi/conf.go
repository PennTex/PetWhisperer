package webapi

import "os"

var (
	imageServiceBasePath     string
	animalServiceBasePath    string
	auth0ClientSecret        string
	servicesAuthorizationKey string
	activityServiceBasePath  string
)

func init() {
	imageServiceBasePath = os.Getenv("IMAGE_SERVICE_BASE_PATH")
	animalServiceBasePath = os.Getenv("ANIMAL_SERVICE_BASE_PATH")
	activityServiceBasePath = os.Getenv("ACTIVITY_SERVICE_BASE_PATH")
	auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	servicesAuthorizationKey = os.Getenv("AUTHORIZATION_KEY")
}
