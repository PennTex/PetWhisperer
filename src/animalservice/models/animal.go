package models

type Animal struct {
	Typ       string   `json:"type"`
	Name      string   `json:"name"`
	Birthday  int64    `json:"birthday"`
	CreatedAt int64    `json:"created_at"`
	Owners    []string `json:"owners"`
	ImageURL  string   `json:"image_url"`
}
