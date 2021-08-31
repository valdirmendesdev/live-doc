package customer

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Repository interface {
	Add(c *entities.Customer) (*entities.Customer, error)
	GetById(id types.ID) (*entities.Customer, error)
	All(limit int, page int) ([]entities.Customer, error)
}
