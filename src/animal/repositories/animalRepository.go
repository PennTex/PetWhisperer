package repositories

import "github.com/PennTex/PetWhisperer/src/animal/models"

type AnimalRepository interface {
	GetAll() []models.Animal
	Get(ID string) *models.Animal
}
