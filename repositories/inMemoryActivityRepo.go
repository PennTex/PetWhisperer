package repositories

import (
	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/PennTex/PetWhisperer/models/animal"
)

var activityDB = createActivities()

type InMemoryActivityRepo struct{}

func createActivities() []activity.Activity {
	animalRepo := InMemoryAnimalRepo{}
	animals := animalRepo.GetAll()
	activities := []activity.Activity{}

	var activity1 = activity.NewFeedActivity(animals[0].GetID(), "mario")
	var activity2 = activity.NewFeedActivity(animals[0].GetID(), "jordan")
	var activity3 = activity.NewMedicationActivity(animals[0].GetID(), "jordan")

	activities = append(activities, activity1, activity2, activity3)

	var fluffy = animal.NewCat("fluffy")
	animals = append(animals, fluffy)

	return activities
}

func (r InMemoryActivityRepo) GetAll() []activity.Activity {
	return activityDB
}

func (r InMemoryActivityRepo) Create(theActivity activity.Activity) error {
	activityDB = append(activityDB, theActivity)
	return nil
}

func (r InMemoryActivityRepo) GetByAnimalID(ID string) []activity.Activity {
	var activities []activity.Activity

	for _, theActivity := range activityDB {
		if theActivity.GetActivityType() == "feed" && theActivity.(*activity.FeedActivity).AnimalID == ID {
			activities = append(activities, theActivity)
		} else if theActivity.GetActivityType() == "medication" && theActivity.(*activity.MedicationActivity).AnimalID == ID {
			activities = append(activities, theActivity)
		}
	}

	return activities
}
