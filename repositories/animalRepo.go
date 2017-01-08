package repositories

import "github.com/PennTex/PetWhisperer/models/animal"

type AnimalRepo interface {
	GetAll() []animal.Animal
	Get(ID string) *animal.Animal
}
