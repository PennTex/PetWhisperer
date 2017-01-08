package repositories

import "github.com/PennTex/PetWhisperer/models/activity"

type ActivityRepo interface {
	GetByAnimalID(ID string) []activity.Activity
}
