package models

import "time"

type Activity struct {
	Typ       string `datastore:"type" json:"type"`
	AnimalID  string `datastore:"animal_id" json:"animal_id"`
	CreatedAt int64  `datastore:"created_at" json:"created_at"`
	By        string `datastore:"by" json:"by"`
	At        string `datastore:"at" json:"at"`
	Note      string `datastore:"note" json:"note"`
}

func NewActivity(typ string, animalID string, by string) *Activity {
	return &Activity{
		Typ:       typ,
		AnimalID:  animalID,
		CreatedAt: time.Now().Unix(),
		By:        by,
	}
}
