package repositories

import "github.com/PennTex/pet-whisperer/src/activity/models"

var activityDB = createActivities()

type InMemoryActivityRepository struct{}

func createActivities() []models.Activity {
	activities := []models.Activity{}

	var activity1 = models.NewFeedActivity("abc", "mario")
	var activity2 = models.NewFeedActivity("abc", "jordan")
	var activity3 = models.NewMedicationActivity("abc", "jordan")

	activities = append(activities, activity1, activity2, activity3)

	return activities
}

func (r InMemoryActivityRepository) Create(theActivity models.Activity) (error, *models.Activity) {
	activityDB = append(activityDB, theActivity)
	return nil, &theActivity
}

func (r InMemoryActivityRepository) GetByAnimalID(ID string) (error, []models.Activity) {
	var activities []models.Activity

	for _, theActivity := range activityDB {
		if theActivity.GetActivityType() == "feed" && theActivity.(*models.FeedActivity).AnimalID == ID {
			activities = append(activities, theActivity)
		} else if theActivity.GetActivityType() == "medication" && theActivity.(*models.MedicationActivity).AnimalID == ID {
			activities = append(activities, theActivity)
		}
	}

	return nil, activities
}
