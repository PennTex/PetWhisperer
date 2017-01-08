package repositories

import "github.com/PennTex/PetWhisperer/models/animal"

// temp fake animal db
var animalDB = []animal.Animal{
	{
		ID:   "5455c3b8-d5da-11e6-bf26-cec0c932ce01",
		Name: "Shawnee",
	},
	{
		ID:   "5455c606-d5da-11e6-bf26-cec0c932ce01",
		Name: "Sheba",
	},
}

type MemoryRepo struct{}

func (r MemoryRepo) GetAll() []animal.Animal {
	return animalDB
}

func (r MemoryRepo) Get(ID string) *animal.Animal {
	for _, animal := range animalDB {
		if animal.ID == ID {
			return &animal
		}
	}
	return nil
}
