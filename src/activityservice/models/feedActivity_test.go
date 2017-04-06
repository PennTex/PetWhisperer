package models_test

import (
	"testing"

	"github.com/PennTex/pet-whisperer/src/activity/models"
	"github.com/stretchr/testify/assert"
)

func TestFeedActivity_ImplementsActivity(t *testing.T) {
	var _ models.Activity = models.FeedActivity{}
}

func TestFeedActivity_New(t *testing.T) {
	activity := models.NewFeedActivity("DOGID", "Mario")

	assert.True(t, activity.CreatedAt > 0)
	assert.True(t, activity.Typ == "feed")
	assert.True(t, activity.FedBy == "Mario")
}
