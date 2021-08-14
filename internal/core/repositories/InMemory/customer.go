package InMemory

import (
	"github.com/valdirmendesdev/live-doc/internal/core/db"
)

type InMemory struct{
	list map[db.ID]*db.Customer
}

func NewInMemory() *InMemory {
	return &InMemory{
		list: map[db.ID]*db.Customer{},
	}
}

func (r *InMemory) Create(e *db.Customer) (*db.Customer, error) {
	r.list[e.ID] = e
	return e, nil
}

