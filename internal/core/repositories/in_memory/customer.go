package in_memory

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type InMemory struct{
	list map[models.ID]*models.Customer
}

func NewInMemory() *InMemory {
	return &InMemory{
		list: map[models.ID]*models.Customer{},
	}
}

func (r *InMemory) Store(e *models.Customer) (*models.Customer, error) {
	r.list[e.ID] = e
	return e, nil
}

