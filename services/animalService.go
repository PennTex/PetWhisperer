package services

import (
	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/PennTex/PetWhisperer/repositories"
)

type AnimalService struct {
	animalRepo repositories.AnimalsRepo
}

func NewAnimalService(animalRepo repositories.AnimalsRepo) *AnimalService {
	return &AnimalService{
		animalRepo: animalRepo,
	}
}

func (a *AnimalService) GetAnimals() []animal.Animal {
	return a.animalRepo.GetAll()
}

func (a *AnimalService) GetAnimalByID(ID string) *animal.Animal {
	return a.animalRepo.Get(ID)
}
