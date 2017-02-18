package app

import "github.com/GoogleCloudPlatform/go-endpoints/endpoints"

func init() {
	// register the quotes API with cloud endpoints.
	api, err := endpoints.RegisterService(AnimalService{}, "animalService", "v1", "Animal API", true)
	if err != nil {
		panic(err)
	}

	// adapt the name, method, and path for each method.
	info := api.MethodByName("List").Info()
	info.Name, info.HTTPMethod, info.Path = "getQuotes", "GET", "quotesService"

	info = api.MethodByName("Add").Info()
	info.Name, info.HTTPMethod, info.Path = "addQuote", "POST", "quotesService"

	// start handling cloud endpoint requests.
	endpoints.HandleHTTP()
}
