package customer

import "github.com/valdirmendesdev/live-doc/entities"

type Repository interface {
	Create(e *entities.Customer) (*entities.Customer, error)
}
