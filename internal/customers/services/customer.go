package services

import (
	"github.com/valdirmendesdev/live-doc/internal/customers/dto"
	"github.com/valdirmendesdev/live-doc/internal/customers/entities"
	"github.com/valdirmendesdev/live-doc/internal/customers/repositories"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Customer struct {
	Repo repositories.Customer
}

func NewCustomer(repo repositories.Customer) *Customer {
	return &Customer{Repo: repo}
}

func (c *Customer) Create(dto dto.CreateRequest) (*entities.Customer, error) {
	customer := entities.NewCustomer()
	customer.FiscalID = dto.FiscalID
	customer.CorporateName = dto.CorporateName
	customer.TradeName = dto.TradeName
	customer.Address = dto.Address
	customer.Number = dto.Number
	customer.City = dto.City
	customer.State = dto.State
	customer.Zip = dto.Zip
	customer.Complement = dto.Complement

	isValid, err := customer.IsValid()
	if !isValid {
		return nil, err
	}

	createdCustomer, err := c.Repo.Create(&customer)
	if err != nil {
		return nil, err
	}

	return createdCustomer, nil
}

func (c *Customer) ListAll(limit int, page int) ([]entities.Customer, error) {
	customersList, err := c.Repo.ListAll(limit, page)
	if err != nil {
		return nil, err
	}
	return customersList, nil
}

func (c *Customer) FindById(id types.ID) (*entities.Customer, error) {
	customer, err := c.Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
