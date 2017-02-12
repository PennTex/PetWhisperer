package repositories

import "github.com/PennTex/PetWhisperer/src/animal/models"

type AnimalRepository interface {
	GetAll() (error, []models.Animal)
	Get(ID string) (error, *models.Animal)
}
