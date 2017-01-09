package activity

import "time"

type ActivityBase struct {
	Typ       string `json:"type"`
	AnimalID  string `json:"animal_id"`
	CreatedAt int64  `json:"created_at"`
}

type Activity interface {
	GetActivityType() string
}

type ActivityPerformer interface {
	GetID() string
}

func newActivity(animalID string, typ string) *ActivityBase {
	return &ActivityBase{
		Typ:       typ,
		AnimalID:  animalID,
		CreatedAt: time.Now().Unix(),
	}
}

func (a *ActivityBase) GetActivityType() string {
	return a.Typ
}
