package animalservice

import (
	"context"
	"time"
)

var animalRepo CloudDatastoreRepository

type AnimalService struct{}

func (s *AnimalService) GetAnimal(ctx context.Context, animalID string) (*Animal, error) {
	animal, err := animalRepo.GetByID(ctx, animalID)

	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (s *AnimalService) GetAnimals(ctx context.Context) ([]Animal, error) {
	animals, err := animalRepo.Get(ctx)

	if err != nil {
		return nil, err
	}

	return animals, nil
}

func (s *AnimalService) CreateAnimal(ctx context.Context, animal *Animal) (string, error) {
	animal.CreatedAt = time.Now().Unix()
	animalID, err := animalRepo.Create(ctx, animal)

	if err != nil {
		return "", err
	}

	return animalID, nil
}
