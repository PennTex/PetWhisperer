package activity

import "time"

type Activity struct {
	Typ       string
	AnimalID  string
	CreatedAt int64
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
