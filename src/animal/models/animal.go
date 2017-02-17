package models

import "time"

type Animal struct {
	Typ       string   `json:"type"`
	Name      string   `json:"name"`
	Birthday  int64    `json:"birthday"`
	CreatedAt int64    `json:"created_at"`
	Owners    []string `json:"owners"`
	ImageURL  string   `json:"image_url"`
}

func newAnimal(typ string, name string) *Animal {
	return &Animal{
		Typ:       typ,
		Name:      name,
		Birthday:  0,
		CreatedAt: time.Now().Unix(),
	}
}
