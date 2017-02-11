package models

type Dog struct {
	*AnimalBase
	Breed []string `json:"breed"`
}

func NewDog(name string) *Dog {
	return &Dog{
		AnimalBase: newAnimal("dog", name),
		Breed:      []string{},
	}
}

func (d *Dog) GetBreed() []string {
	return d.Breed
}

func (d *Dog) SetBreed(breed []string) {
	d.Breed = breed
}
