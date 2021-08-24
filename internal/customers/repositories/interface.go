package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/customers/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Customer interface {
	Create(c *entities.Customer) (*entities.Customer, error)
	FindById(id types.ID) (*entities.Customer, error)
	ListAll(limit int, page int) ([]entities.Customer, error)
}
