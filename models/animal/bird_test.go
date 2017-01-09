package animal_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/stretchr/testify/assert"
)

func TestBird_ImplementsAnimal(t *testing.T) {
	var _ animal.Animal = animal.Bird{}
}

func TestBird_New(t *testing.T) {
	birdName := "Fluffy"
	tweety := animal.NewBird(birdName)

	assert.Equal(t, tweety.Name, birdName)
	assert.Equal(t, tweety.Typ, "bird")
	assert.True(t, tweety.CreatedAt > 0)
}
