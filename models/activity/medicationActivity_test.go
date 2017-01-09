package activity_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/stretchr/testify/assert"
)

func TestMedicationActivity_ImplementsActivity(t *testing.T) {
	var _ activity.Activity = activity.MedicationActivity{}
}

func TestMedicationActivity_New(t *testing.T) {
	activity := activity.NewMedicationActivity("DOGID", "Mario")

	assert.True(t, activity.CreatedAt > 0)
	assert.True(t, activity.Typ == "medication")
	assert.True(t, activity.GivenBy == "Mario")
}
