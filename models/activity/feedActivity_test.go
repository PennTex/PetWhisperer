package activity_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/activity"
	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/stretchr/testify/assert"
)

func TestFeedActivity_New(t *testing.T) {
	shawnee := animal.NewDog("shawnee")
	activity := activity.NewFeedActivity(shawnee, "Mario")

	assert.True(t, activity.CreatedAt > 0)
	assert.True(t, activity.Typ == "feed")
	assert.True(t, activity.FedBy == "Mario")
}
