package customer

import (
	"github.com/valdirmendesdev/live-doc/entities"
)

type InMemory struct{
	list map[entities.ID]*entities.Customer
}

func NewInMemory() *InMemory {
	return &InMemory{
		list: map[entities.ID]*entities.Customer{},
	}
}

func (r *InMemory) Create(e *entities.Customer) (*entities.Customer, error) {
	r.list[e.ID] = e
	return e, nil
}

