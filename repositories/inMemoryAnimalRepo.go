package repositories

import "github.com/PennTex/PetWhisperer/models/animal"

var animalDB = []animal.Animal{
	{
		ID:   "5455c3b8-d5da-11e6-bf26-cec0c932ce01",
		Typ:  "dog",
		Name: "Shawnee",
	},
	{
		ID:   "5455c606-d5da-11e6-bf26-cec0c932ce01",
		Typ:  "dog",
		Name: "Sheba",
	},
	{
		ID:   "5456c603-d5da-11e5-bf26-cec0c932ce01",
		Typ:  "cat",
		Name: "Matt",
	},
}

type InMemoryAnimalRepo struct{}

func (r InMemoryAnimalRepo) GetAll() []animal.Animal {
	return animalDB
}

func (r InMemoryAnimalRepo) Get(ID string) *animal.Animal {
	for _, animal := range animalDB {
		if animal.ID == ID {
			return &animal
		}
	}
	return nil
}
