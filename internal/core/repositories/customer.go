package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type Customer interface {
	Create(c *models.Customer) (*models.Customer, error)
	FindById(id models.ID) (*models.Customer, error)
	ListAll(limit int, page int) ([]models.Customer, error)
}
