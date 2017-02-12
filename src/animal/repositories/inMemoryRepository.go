package repositories

import "github.com/PennTex/PetWhisperer/src/animal/models"

var animalDB = createAnimals()

type InMemoryAnimalRepository struct{}

func createAnimals() []models.Animal {
	animals := []models.Animal{}

	var shawnee = models.NewDog("shawnee")
	shawnee.SetBreed([]string{"border collie", "lab"})
	animals = append(animals, shawnee)

	var fluffy = models.NewCat("fluffy")
	animals = append(animals, fluffy)

	return animals
}

func (r InMemoryAnimalRepository) GetAll() []models.Animal {
	return animalDB
}

func (r InMemoryAnimalRepository) Get(ID string) *models.Animal {
	for _, animal := range r.GetAll() {
		if animal.GetID() == ID {
			return &animal
		}
	}
	return nil
}
