package repositories

import (
	"golang.org/x/net/context"

	"github.com/PennTex/pet-whisperer/src/activityservice/models"
)

type ActivityRepository interface {
	Create(ctx context.Context, animal *models.Activity) (error, *models.Activity)
	GetByAnimalID(ctx context.Context, ID string) (error, []models.Activity)
}
