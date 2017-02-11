package models

type MedicationActivity struct {
	*ActivityBase
	GivenBy string `json:"given_by"`
}

func NewMedicationActivity(animalID string, givenBy string) *MedicationActivity {
	return &MedicationActivity{
		ActivityBase: newActivity(animalID, "medication"),
		GivenBy:      givenBy,
	}
}
