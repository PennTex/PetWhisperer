package repositories

import "github.com/PennTex/PetWhisperer/src/activity/models"

type ActivityRepository interface {
	Create(theActivity models.Activity) error
	GetByAnimalID(ID string) []models.Activity
}
