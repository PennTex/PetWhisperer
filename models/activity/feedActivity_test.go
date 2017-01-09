package activity_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/stretchr/testify/assert"
)

func TestFeedActivity_ImplementsActivity(t *testing.T) {
	var _ activity.Activity = activity.FeedActivity{}
}

func TestFeedActivity_New(t *testing.T) {
	activity := activity.NewFeedActivity("DOGID", "Mario")

	assert.True(t, activity.CreatedAt > 0)
	assert.True(t, activity.Typ == "feed")
	assert.True(t, activity.FedBy == "Mario")
}
