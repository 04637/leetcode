package q4

import (
	"fmt"
	"leetcode/learn/chapter1/lib"
)

type PetQueue struct {
	dogQ  lib.Queue
	catQ  lib.Queue
	count int
}

func (p *PetQueue) Add(pet interface{}) {

	switch pet.(type) {
	case *Dog:
		ptq := new(PetToQueue)
		ptq.data = pet
		ptq.count = p.count
		p.count++
		p.dogQ.Add(*ptq)
	case *Cat:
		ptq := new(PetToQueue)
		ptq.data = pet
		ptq.count = p.count
		p.count++
		p.catQ.Add(*ptq)
	default:
		fmt.Println("type not found")
	}
}

func (p *PetQueue) PollAll() interface{} {
	pet1 := p.catQ.Peek()
	pet2 := p.dogQ.Peek()
	if pet1 != nil && pet2 != nil {
		cat := pet1.(PetToQueue)
		dog := pet2.(PetToQueue)
		if cat.count < dog.count {
			return p.catQ.Poll().(PetToQueue).data
		} else {
			return p.dogQ.Poll().(PetToQueue).data
		}
	} else {
		if pet1 == nil {
			return p.dogQ.Poll().(PetToQueue).data
		} else {
			return p.catQ.Poll().(PetToQueue).data
		}
	}
}

func (p *PetQueue) PollDog() *Dog {
	return p.dogQ.Poll().(PetToQueue).data.(*Dog)
}

func (p *PetQueue) PollCat() *Cat {
	return p.catQ.Poll().(PetToQueue).data.(*Cat)
}

func (p *PetQueue) IsEmpty() bool {
	return p.dogQ.IsEmpty() && p.catQ.IsEmpty()
}

func (p *PetQueue) IsDogEmpty() bool {
	return p.dogQ.IsEmpty()
}

func (p *PetQueue) IsCatEmpty() bool {
	return p.catQ.IsEmpty()
}

func (p *PetQueue) Size() int {
	return p.catQ.Size() + p.dogQ.Size()
}

type PetToQueue struct {
	data  interface{}
	count int
}
