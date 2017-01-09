package repositories

import "github.com/PennTex/PetWhisperer/models/activity"

type ActivityRepo interface {
	Create(theActivity activity.Activity) error
	GetByAnimalID(ID string) []activity.Activity
}
