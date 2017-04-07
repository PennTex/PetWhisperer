package repositories

import (
	"golang.org/x/net/context"

	"github.com/PennTex/pet-whisperer/src/activityservice/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type CloudDatastoreRepository struct{}

var entityKind = "activity"

func (r CloudDatastoreRepository) Create(ctx context.Context, activity *models.Activity) (*models.Activity, error) {
	key := datastore.NewKey(ctx, entityKind, uuid.NewV4().String(), 0, nil)

	_, err := datastore.Put(ctx, key, activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (r CloudDatastoreRepository) GetByAnimalID(ctx context.Context, ID string) ([]models.Activity, error) {
	var activities []models.Activity
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
