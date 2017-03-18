package models

type Animal struct {
	ID        string   `datastore:"-" json:"id"`
	Typ       string   `datastore:"type" json:"type"`
	Name      string   `datastore:"name" json:"name"`
	Birthday  int64    `datastore:"birthday" json:"birthday"`
	CreatedAt int64    `datastore:"created_at" json:"created_at"`
	Owners    []string `datastore:"owners" json:"owners"`
	ImageURL  string   `datastore:"image_url" json:"image_url"`
}
