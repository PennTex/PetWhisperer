package repositories

import (
	"log"

	"golang.org/x/net/context"

	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
)

type CloudDatastoreRepository struct{}

var entityKind = "animal"

func (r CloudDatastoreRepository) Create(ctx context.Context, animal *models.Animal) (*models.Animal, error) {
	key := datastore.NewKey(ctx, entityKind, uuid.NewV4().String(), 0, nil)

	animalKey, err := datastore.Put(ctx, key, animal)
	if err != nil {
		return nil, err
	}

	animal.ID = animalKey.StringID()

	return animal, nil
}

func (r CloudDatastoreRepository) Get(ctx context.Context) ([]models.Animal, error) {
	return nil, nil
}

func (r CloudDatastoreRepository) GetByOwnerID(ctx context.Context, ID string) ([]models.Animal, error) {
	var animals []models.Animal

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

func (r CloudDatastoreRepository) GetByID(ctx context.Context, ID string) (*models.Animal, error) {
	var animal models.Animal

	animalKey := datastore.NewKey(ctx, entityKind, ID, 0, nil)

	if err := datastore.Get(ctx, animalKey, &animal); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &animal, nil
}

func (r CloudDatastoreRepository) Destroy(ctx context.Context, ID string) error {
	animalKey := datastore.NewKey(ctx, entityKind, ID, 0, nil)
	return datastore.Delete(ctx, animalKey)
}
