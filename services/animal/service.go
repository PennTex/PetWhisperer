package main

import (
	"github.com/PennTex/PetWhisperer/services/animal/models"
	"github.com/PennTex/PetWhisperer/services/animal/repositories"
)

type AnimalService struct {
	animalRepository repositories.AnimalRepository
}

func NewAnimalService(animalRepository repositories.AnimalRepository) *AnimalService {
	return &AnimalService{
		animalRepository: animalRepository,
	}
}

func (a *AnimalService) GetAnimals() []models.Animal {
	return a.animalRepository.GetAll()
}

func (a *AnimalService) GetAnimal(ID string) *models.Animal {
	return a.animalRepository.Get(ID)
}
