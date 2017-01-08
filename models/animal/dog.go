package animal

type Dog struct {
	*Animal
	breed []string
}

func NewDog(name string) *Dog {
	return &Dog{
		Animal: newAnimal("dog", name),
	}
}

func (d *Dog) GetBreed() []string {
	return d.breed
}

func (d *Dog) SetBreed(breed []string) {
	d.breed = breed
}
