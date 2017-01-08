package services

import (
	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/PennTex/PetWhisperer/repositories"
)

type AnimalService struct {
	animalRepo repositories.AnimalRepo
}

func NewAnimalService(animalRepo repositories.AnimalRepo) *AnimalService {
	return &AnimalService{
		animalRepo: animalRepo,
	}
}

func (a *AnimalService) GetAnimals() []animal.Animal {
	return a.animalRepo.GetAll()
}

func (a *AnimalService) GetAnimal(ID string) *animal.Animal {
	return a.animalRepo.Get(ID)
}
