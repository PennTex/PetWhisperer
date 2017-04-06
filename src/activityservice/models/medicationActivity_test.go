package models_test

import (
	"testing"

	"github.com/PennTex/pet-whisperer/src/activity/models"
	"github.com/stretchr/testify/assert"
)

func TestMedicationActivity_ImplementsActivity(t *testing.T) {
	var _ models.Activity = models.MedicationActivity{}
}

func TestMedicationActivity_New(t *testing.T) {
	activity := models.NewMedicationActivity("DOGID", "Mario")

	assert.True(t, activity.CreatedAt > 0)
	assert.True(t, activity.Typ == "medication")
	assert.True(t, activity.GivenBy == "Mario")
}
