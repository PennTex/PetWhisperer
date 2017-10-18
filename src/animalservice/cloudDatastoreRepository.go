package animalservice

import (
	"log"

	"golang.org/x/net/context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
)

type animal struct {
	ID        string   `datastore:"-" json:"id"`
	Typ       string   `datastore:"type" json:"type"`
	Name      string   `datastore:"name" json:"name"`
	Birthday  int64    `datastore:"birthday" json:"birthday"`
	CreatedAt int64    `datastore:"created_at" json:"created_at"`
	Owners    []string `datastore:"owners" json:"owners"`
	ImageURL  string   `datastore:"image_url" json:"image_url"`
}

type cloudDatastoreRepository struct{}

var entityKind = "animal"

func (r cloudDatastoreRepository) create(ctx context.Context, animal *animal) (*animal, error) {
	key := datastore.NewKey(ctx, entityKind, uuid.NewV4().String(), 0, nil)

	animalKey, err := datastore.Put(ctx, key, animal)
	if err != nil {
		return nil, err
	}

	animal.ID = animalKey.StringID()

	return animal, nil
}

func (r cloudDatastoreRepository) get(ctx context.Context) ([]animal, error) {
	return nil, nil
}

func (r cloudDatastoreRepository) getByOwnerID(ctx context.Context, ID string) ([]animal, error) {
	var animals []animal

	query := datastore.
		NewQuery(entityKind).
		Filter("owners =", ID)

	keys, err := query.GetAll(ctx, &animals)
	if err != nil {
		return nil, err
	}

	for i := range animals {
		animals[i].ID = keys[i].StringID()
	}

	log.Printf("returned: %s", keys)

	return animals, nil
}

func (r cloudDatastoreRepository) getByID(ctx context.Context, ID string) (*animal, error) {
	var animal animal

	animalKey := datastore.NewKey(ctx, entityKind, ID, 0, nil)

	if err := datastore.Get(ctx, animalKey, &animal); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &animal, nil
}

func (r cloudDatastoreRepository) destroy(ctx context.Context, ID string) error {
	animalKey := datastore.NewKey(ctx, entityKind, ID, 0, nil)
	return datastore.Delete(ctx, animalKey)
}
