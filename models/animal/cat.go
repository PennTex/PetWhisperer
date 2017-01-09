package animal

type Cat struct {
	*AnimalBase
}

func NewCat(name string) *Cat {
	return &Cat{
		AnimalBase: newAnimal("cat", name),
	}
}
