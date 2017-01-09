package services

import (
	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/PennTex/PetWhisperer/repositories"
)

type ActivityService struct {
	activityRepo repositories.ActivityRepo
}

func NewActivityService(activityRepo repositories.ActivityRepo) *ActivityService {
	return &ActivityService{
		activityRepo: activityRepo,
	}
}

func (a *ActivityService) GetAnimalActivity(animalID string) []activity.Activity {
	return a.activityRepo.GetByAnimalID(animalID)
}

func (a *ActivityService) CreateActivity(theActivity activity.Activity) error {
	return a.activityRepo.Create(theActivity)
}
