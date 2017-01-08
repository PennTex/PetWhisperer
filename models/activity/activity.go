package activity

import "time"

type Activity struct {
	Typ       string `json:"type"`
	AnimalID  string `json:"animal_id"`
	CreatedAt int64  `json:"created_at"`
}

type ActivityPerformer interface {
	GetID() string
}

func newActivity(typ string, performedBy ActivityPerformer) *Activity {
	return &Activity{
		Typ:       typ,
		AnimalID:  performedBy.GetID(),
		CreatedAt: time.Now().Unix(),
	}
}
