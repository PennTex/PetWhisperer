package animal_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/models/animal"
	"github.com/stretchr/testify/assert"
)

func TestCat_ImplementsAnimal(t *testing.T) {
	var _ animal.Animal = animal.Cat{}
}

func TestCat_New(t *testing.T) {
	catName := "Fluffy"
	fluffy := animal.NewCat(catName)

	assert.Equal(t, fluffy.Name, catName)
	assert.Equal(t, fluffy.Typ, "cat")
	assert.True(t, fluffy.CreatedAt > 0)
}
