package models_test

import (
	"testing"

	"github.com/PennTex/PetWhisperer/src/animal/models"
	"github.com/stretchr/testify/assert"
)

func TestCat_ImplementsAnimal(t *testing.T) {
	var _ models.Animal = models.Cat{}
}

func TestCat_New(t *testing.T) {
	catName := "Fluffy"
	fluffy := models.NewCat(catName)

	assert.Equal(t, fluffy.Name, catName)
	assert.Equal(t, fluffy.Typ, "cat")
	assert.True(t, fluffy.CreatedAt > 0)
}
