package app

import (
	"encoding/json"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile := session.Values["profile"].(map[string]interface{})
	nickname := profile["nickname"].(string)
	appMetadata := profile["app_metadata"].(map[string]interface{})
	pets := appMetadata["pets"].([]interface{})

	data := struct {
		Nickname string
		Pets     []interface{}
	}{
		Nickname: nickname,
		Pets:     pets,
	}

	RenderTemplate(w, "dashboard", data)
}

func getAppMetadataFromProfile(profile interface{}) (map[string]interface{}, error) {
	var app_metadata map[string]interface{}

	if err := json.Unmarshal([]byte(profile.(string)), &app_metadata); err != nil {
		return nil, err
	}

	return app_metadata, nil
}
