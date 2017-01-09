package repositories

import "github.com/PennTex/PetWhisperer/models/animal"

var animalDB = createAnimals()

type InMemoryAnimalRepo struct{}

func createAnimals() []animal.Animal {
	animals := []animal.Animal{}

	var shawnee = animal.NewDog("shawnee")
	shawnee.SetBreed([]string{"border collie", "lab"})
	animals = append(animals, shawnee)

	var fluffy = animal.NewCat("fluffy")
	animals = append(animals, fluffy)

	return animals
}

func (r InMemoryAnimalRepo) GetAll() []animal.Animal {
	return animalDB
}

func (r InMemoryAnimalRepo) Get(ID string) *animal.Animal {
	for _, animal := range r.GetAll() {
		if animal.GetID() == ID {
			return &animal
		}
	}
	return nil
}
