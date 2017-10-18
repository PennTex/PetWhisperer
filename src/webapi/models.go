package webapi

type animalPostReq struct {
	Typ      string   `json:"type"`
	Name     string   `json:"name"`
	Birthday int64    `json:"birthday"`
	Owners   []string `json:"owners"`
	ImageURL string   `json:"image_url"`
}

type activityPostReq struct {
	Typ  string `json:"type"`
	By   string `json:"by"`
	At   int64  `json:"at"`
	Note string `json:"note"`
}
