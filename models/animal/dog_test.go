package animal_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/stretchr/testify/assert"
)

func TestDog_New(t *testing.T) {
	dogName := "Max"
	max := animal.NewDog(dogName)

	assert.Equal(t, max.Name, dogName)
	assert.Equal(t, max.Typ, "dog")
	assert.True(t, max.CreatedAt > 0)
}

func TestDog_GetBreed(t *testing.T) {
	max := animal.NewDog("Max")

	breed := max.GetBreed()

	assert.True(t, len(breed) == 0)
}

func TestDog_SetBreed(t *testing.T) {
	max := animal.NewDog("Max")
	breed := []string{"lab", "pit"}
	max.SetBreed(breed)

	assert.Equal(t, max.GetBreed(), breed)
}
