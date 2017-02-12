package models_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/src/animal/models"
	"github.com/stretchr/testify/assert"
)

func TestBird_ImplementsAnimal(t *testing.T) {
	var _ models.Animal = models.Bird{}
}

func TestBird_New(t *testing.T) {
	birdName := "Fluffy"
	tweety := models.NewBird(birdName)

	assert.Equal(t, tweety.Name, birdName)
	assert.Equal(t, tweety.Typ, "bird")
	assert.True(t, tweety.CreatedAt > 0)
}
