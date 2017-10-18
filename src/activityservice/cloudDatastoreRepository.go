package activityservice

import (
	"golang.org/x/net/context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type activity struct {
	Typ       string `datastore:"type" json:"type"`
	AnimalID  string `datastore:"animal_id" json:"animal_id"`
	CreatedAt int64  `datastore:"created_at" json:"created_at"`
	By        string `datastore:"by" json:"by"`
	At        int64  `datastore:"at" json:"at"`
	Note      string `datastore:"note" json:"note"`
}

type cloudDatastoreRepository struct{}

var entityKind = "activity"

func (r cloudDatastoreRepository) create(ctx context.Context, a *activity) (*activity, error) {
	key := datastore.NewKey(ctx, entityKind, uuid.NewV4().String(), 0, nil)

	_, err := datastore.Put(ctx, key, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r cloudDatastoreRepository) getByAnimalID(ctx context.Context, ID string) ([]activity, error) {
	var activities []activity
	log.Infof(ctx, "Getting Activities for : %s", ID)

	query := datastore.
		NewQuery(entityKind).
		Filter("animal_id =", ID)

	_, err := query.GetAll(ctx, &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}
