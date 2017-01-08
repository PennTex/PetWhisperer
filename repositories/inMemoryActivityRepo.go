package repositories

import "github.com/PennTex/PetWhisperer/models/activity"

var activityDB = []activity.Activity{
	{
		AnimalID: "5455c3b8-d5da-11e6-bf26-cec0c932ce01",
		Typ:      "feed",
	},
	{
		AnimalID: "5455c606-d5da-11e6-bf26-cec0c932ce01",
		Typ:      "feed",
	},
}

type InMemoryActivityRepo struct{}

func (r InMemoryActivityRepo) GetByAnimalID(ID string) []activity.Activity {
	var activities []activity.Activity

	for _, activity := range activityDB {
		if activity.AnimalID == ID {
			activities = append(activities, activity)
		}
	}

	return activities
}
