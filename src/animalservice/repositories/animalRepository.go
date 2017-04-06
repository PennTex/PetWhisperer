package repositories

import (
	"golang.org/x/net/context"

	"github.com/PennTex/pet-whisperer/src/animalservice/models"
)

type AnimalRepository interface {
	Create(ctx context.Context, animal *models.Animal) (*models.Animal, error)
	Get(ctx context.Context) ([]models.Animal, error)
	GetByID(ctx context.Context, ID string) (*models.Animal, error)
	GetByOwnerID(ctx context.Context, ID string) ([]models.Animal, error)
	Destroy(ctx context.Context, ID string) error
}
