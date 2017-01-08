package activity

import "time"

type activity struct {
	AnimalID  string
	CreatedAt int64
}

type ActivityPerformer interface {
	GetID() string
}

func newActivity(performedBy ActivityPerformer) *activity {
	return &activity{
		AnimalID:  performedBy.GetID(),
		CreatedAt: time.Now().Unix(),
	}
}
