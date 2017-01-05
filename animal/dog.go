package animal

type Dog struct {
	*animal
	breed []string
}

func NewDog(name string) *Dog {
	return &Dog{
		animal: newAnimal("dog", name),
	}
}

func (d *Dog) GetBreed() []string {
	return d.breed
}

func (d *Dog) SetBreed(breed []string) {
	d.breed = breed
}
