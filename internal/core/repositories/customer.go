package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type UserRepository interface {
	Store(e *models.Customer) (*models.Customer, error)
}
