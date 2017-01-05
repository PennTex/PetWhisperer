package animal

import (
	"time"
)

type animal struct {
	Typ       string
	Name      string
	Birthday  int64
	CreatedAt int64
}

func newAnimal(typ string, name string) *animal {
	return &animal{
		Typ:       typ,
		Name:      name,
		Birthday:  0,
		CreatedAt: time.Now().Unix(),
	}
}
