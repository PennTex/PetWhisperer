package repositories

import "github.com/PennTex/PetWhisperer/models/animal"

type AnimalsRepo interface {
	GetAll() []animal.Animal
	Get(ID string) *animal.Animal
}
