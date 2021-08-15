package customer

import "github.com/valdirmendesdev/live-doc/internal/core/models"

type CustomerDTO struct {
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

type UseCase interface {
	CreateCustomer(customerData CustomerDTO) (*models.Customer, error)
}
