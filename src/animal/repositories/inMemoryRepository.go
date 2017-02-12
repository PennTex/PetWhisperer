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

func (r InMemoryAnimalRepository) GetAll() (error, []models.Animal) {
	return nil, animalDB
}

func (r InMemoryAnimalRepository) Get(ID string) (error, *models.Animal) {
	_, animals := r.GetAll()

	for _, animal := range animals {
		if animal.GetID() == ID {
			return nil, &animal
		}
	}

	return nil, nil
}
