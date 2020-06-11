package q4

type Pet struct {
	Type string
	Name string
}

func NewPet(class string) *Pet {
	return &Pet{Type: class}
}

type Dog struct {
	Pet
}

func NewDog(name string) *Dog {
	dog := new(Dog)
	dog.Type = "dog"
	dog.Name = name
	return dog
}

type Cat struct {
	Pet
}

func NewCat(name string) *Cat {
	cat := new(Cat)
	cat.Type = "cat"
	cat.Name = name
	return cat
}
