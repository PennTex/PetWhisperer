package repositories

import (
	"log"

	"golang.org/x/net/context"

	"github.com/PennTex/PetWhisperer/src/animalservice/models"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine/datastore"
)

type CloudDatastoreRepository struct{}

func (r CloudDatastoreRepository) Create(ctx context.Context, animal *models.Animal) (string, error) {
	key := datastore.NewKey(ctx, "Animal", uuid.NewV4().String(), 0, nil)
	animalKey, err := datastore.Put(ctx, key, animal)
	if err != nil {
		return "", err
	}

	return animalKey.StringID(), nil
}

func (r CloudDatastoreRepository) Get(ctx context.Context) ([]models.Animal, error) {
	return nil, nil
}

func (r CloudDatastoreRepository) GetByOwnerID(ctx context.Context, ID string) ([]models.Animal, error) {
	query := datastore.NewQuery("Animal").
		Filter("Owners =", ID)

	var animals []models.Animal
	if _, err := query.GetAll(ctx, &animals); err != nil {
		return nil, err
	}

	return animals, nil
}

func (r CloudDatastoreRepository) GetByID(ctx context.Context, ID string) (*models.Animal, error) {
	var animal models.Animal
	animalKey := datastore.NewKey(ctx, "Animal", ID, 0, nil)

	if err := datastore.Get(ctx, animalKey, &animal); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &animal, nil
}
