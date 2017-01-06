package animal

import (
	"time"
)

type Animal struct {
	ID        string
	Name      string
	Birthday  int64
	CreatedAt int64
}

func newAnimal(name string) *Animal {
	return &Animal{
		Name:      name,
		Birthday:  0,
		CreatedAt: time.Now().Unix(),
	}
}

func (a *Animal) GetID() string {
	return a.ID
}
