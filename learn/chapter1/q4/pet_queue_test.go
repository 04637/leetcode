package q4

import (
	"strconv"
	"testing"
)

var petQ *PetQueue

func init() {
	petQ = new(PetQueue)
}

func TestPetQueue_Add(t *testing.T) {
	for i := 0; i < 5; i++ {
		cat := NewCat("cat" + strconv.Itoa(i))
		dog := NewDog("dog" + strconv.Itoa(i))
		petQ.Add(cat)
		petQ.Add(dog)
	}
	if petQ.Size() != 10 {
		t.Errorf("queue size is not correct, "+
			"expected 10 and got %d", petQ.Size())
	}
}

func TestPetQueue_PollAll(t *testing.T) {
	pet2 := petQ.PollAll().(*Cat)
	if pet2.Name != "cat0" {
		t.Errorf("queue poll all error, expected cat0 but got %s", pet2.Name)
	}
	pet := petQ.PollAll().(*Dog)
	if pet.Name != "dog0" {
		t.Errorf("queue poll all error, expected dog0 but got %s", pet.Name)
	}
}

func TestPetQueue_PollCat(t *testing.T) {
	pet := petQ.PollCat()
	if pet.Name != "cat1" {
		t.Errorf("queue poll cat error, expected cat1 but got %s", pet.Name)
	}
}

func TestPetQueue_PollDog(t *testing.T) {
	pet := petQ.PollDog()
	if pet.Name != "dog1" {
		t.Errorf("queue poll dog error, expected dog1 but got %s", pet.Name)
	}
}
