package api

type AnimalPostReq struct {
	Typ      string   `json:"type"`
	Name     string   `json:"name"`
	Birthday int64    `json:"birthday"`
	Owners   []string `json:"owners"`
	ImageURL string   `json:"image_url"`
}

type ActivityPostReq struct {
	Typ  string `json:"type"`
	By   string `json:"by"`
	At   string `json:"at"`
	Note string `json:"note"`
}
