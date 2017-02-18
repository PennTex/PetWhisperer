package repositories

import (
	"golang.org/x/net/context"

	"github.com/PennTex/PetWhisperer/src/animalservice/models"
)

type AnimalRepository interface {
	Create(ctx context.Context, animal *models.Animal) (string, error)
	Get(ctx context.Context) ([]models.Animal, error)
	GetByID(ctx context.Context, ID string) (*models.Animal, error)
}
