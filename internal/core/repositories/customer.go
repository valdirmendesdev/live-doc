package repositories

import (
	"github.com/valdirmendesdev/live-doc/internal/core/db"
)

type UserRepository interface {
	Create(e *db.Customer) (*db.Customer, error)
}
