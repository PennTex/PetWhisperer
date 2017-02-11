package main

import (
	"github.com/PennTex/PetWhisperer/services/activity/models"
	"github.com/PennTex/PetWhisperer/services/activity/repositories"
)

type ActivitiesService struct {
	activityRepository repositories.ActivityRepository
}

func NewActivitiesService(activityRepository repositories.ActivityRepository) *ActivitiesService {
	return &ActivitiesService{
		activityRepository: activityRepository,
	}
}

func (a *ActivitiesService) GetActivitiesByAnimalID(animalID string) []models.Activity {
	return a.activityRepository.GetByAnimalID(animalID)
}

func (a *ActivitiesService) CreateActivity(theActivity models.Activity) error {
	return a.activityRepository.Create(theActivity)
}
