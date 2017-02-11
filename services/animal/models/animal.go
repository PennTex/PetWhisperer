package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type AnimalBase struct {
	ID        string `json:"id"`
	Typ       string `json:"type"`
	Name      string `json:"name"`
	Birthday  int64  `json:"birthday"`
	CreatedAt int64  `json:"created_at"`
}

type Animal interface {
	GetID() string
}

func newAnimal(typ string, name string) *AnimalBase {
	return &AnimalBase{
		ID:        uuid.NewV4().String(),
		Typ:       typ,
		Name:      name,
		Birthday:  0,
		CreatedAt: time.Now().Unix(),
	}
}

func (a *AnimalBase) GetID() string {
	return a.ID
}
