package animal

import (
	"time"
)

type Animal struct {
	ID        string `json:"id"`
	Typ       string `json:"type"`
	Name      string `json:"name"`
	Birthday  int64  `json:"birthday"`
	CreatedAt int64  `json:"created_at"`
}

func newAnimal(typ string, name string) *Animal {
	return &Animal{
		Typ:       typ,
		Name:      name,
		Birthday:  0,
		CreatedAt: time.Now().Unix(),
	}
}

func (a *Animal) GetID() string {
	return a.ID
}
