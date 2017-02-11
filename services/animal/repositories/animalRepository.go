package repositories

import "github.com/PennTex/PetWhisperer/services/animal/models"

type AnimalRepository interface {
	GetAll() []models.Animal
	Get(ID string) *models.Animal
}
