package animalservice

import (
	"time"

	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	"github.com/PennTex/PetWhisperer/src/animalservice/repositories"

	"golang.org/x/net/context"
)

type AnimalService struct {
	animalRepo repositories.AnimalRepository
}

func NewAnimalService(animalRepo repositories.AnimalRepository) *AnimalService {
	return &AnimalService{
		animalRepo: animalRepo,
	}
}

func (s *AnimalService) GetAnimal(ctx context.Context, animalID string) (*models.Animal, error) {
	animal, err := s.animalRepo.GetByID(ctx, animalID)

	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (s *AnimalService) GetAnimals(ctx context.Context) ([]models.Animal, error) {
	animals, err := s.animalRepo.Get(ctx)

	if err != nil {
		return nil, err
	}

	return animals, nil
}

func (s *AnimalService) GetAnimalsByOwnerID(ctx context.Context, ID string) ([]models.Animal, error) {
	animals, err := s.animalRepo.GetByOwnerID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return animals, nil
}

func (s *AnimalService) CreateAnimal(ctx context.Context, animal *models.Animal) (*models.Animal, error) {
	animal.CreatedAt = time.Now().Unix()
	return s.animalRepo.Create(ctx, animal)
}

func (s *AnimalService) DeleteAnimal(ctx context.Context, ID string) error {
	return s.animalRepo.Destroy(ctx, ID)
}
