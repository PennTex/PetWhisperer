package animalservice

import (
	"log"

	"golang.org/x/net/context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
)

type CloudDatastoreRepository struct{}

func (r CloudDatastoreRepository) Create(ctx context.Context, animal *Animal) (string, error) {
	key := datastore.NewKey(ctx, "Animal", uuid.NewV4().String(), 0, nil)
	animalKey, err := datastore.Put(ctx, key, animal)
	if err != nil {
		return "", err
	}

	return animalKey.StringID(), nil
}

func (r CloudDatastoreRepository) Get(ctx context.Context) ([]Animal, error) {
	return nil, nil
}

func (r CloudDatastoreRepository) GetByID(ctx context.Context, ID string) (*Animal, error) {
	var animal Animal
	animalKey := datastore.NewKey(ctx, "Animal", ID, 0, nil)

	if err := datastore.Get(ctx, animalKey, &animal); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &animal, nil
}
