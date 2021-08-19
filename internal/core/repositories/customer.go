package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type Customer interface {
	Create(e *models.Customer) (*models.Customer, error)
}
