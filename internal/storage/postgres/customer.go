package postgres

import (
	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/config"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type CustomerRepository struct{}

func NewCustomerRepository() CustomerRepository {
	return CustomerRepository{}
}

func (r *CustomerRepository) Add(c *entities.Customer) (*entities.Customer, error) {
	err := config.DB.Create(c).Error
	return c, err
}

func (r *CustomerRepository) GetById(id types.ID) (*entities.Customer, error) {
	c := new(entities.Customer)
	err := config.DB.Where(`id = ?`, id).First(c).Error
	return c, err
}

func (r *CustomerRepository) All(limit int, page int) ([]entities.Customer, error) {
	var l []entities.Customer
	err := config.DB.Find(&l).Error
	return l, err
}
