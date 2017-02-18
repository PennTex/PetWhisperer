package animalservice

import "context"

type Animal struct {
	Typ       string   `json:"type"`
	Name      string   `json:"name"`
	Birthday  int64    `json:"birthday"`
	CreatedAt int64    `json:"created_at"`
	Owners    []string `json:"owners"`
	ImageURL  string   `json:"image_url"`
}

type AnimalRepository interface {
	Create(ctx context.Context, animal *Animal) (string, error)
	Get(ctx context.Context) ([]Animal, error)
	GetByID(ctx context.Context, ID string) (*Animal, error)
}
