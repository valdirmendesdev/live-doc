package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type Customer interface {
	Create(c *models.Customer) (*models.Customer, error)
	List() ([]models.Customer, error)
}
