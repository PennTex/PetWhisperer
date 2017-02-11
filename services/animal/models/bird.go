package models

type Bird struct {
	*AnimalBase
}

func NewBird(name string) *Bird {
	return &Bird{
		AnimalBase: newAnimal("bird", name),
	}
}
