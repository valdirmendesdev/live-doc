package usecases

import (
	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"github.com/valdirmendesdev/live-doc/internal/core/repositories"
)

type Customer struct {
	Repo repositories.Customer
}

func NewCustomer(repo repositories.Customer) *Customer {
	return &Customer{Repo: repo}
}

type CustomerCreateRequest struct {
	FiscalID      string
	CorporateName string
	TradeName     string
	Address       string
	Number        string
	City          string
	State         string
	Zip           string
	Complement    string
}

func (c *Customer) Create(dto CustomerCreateRequest) (*models.Customer, error) {
	customer := models.NewCustomer()
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
