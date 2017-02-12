package repositories

import "github.com/PennTex/PetWhisperer/src/activity/models"

type ActivityRepository interface {
	Create(theActivity models.Activity) (error, *models.Activity)
	GetByAnimalID(ID string) (error, []models.Activity)
}
