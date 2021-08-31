package customer

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
)

type CreateDTO struct {
	FiscalID      string `json:"fiscal_id"`
	CorporateName string `json:"corporate_name"`
	TradeName     string `json:"trade_name"`
	Address       string `json:"address"`
	Number        string `json:"number"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Complement    string `json:"complement"`
}

type CreateService struct {
	Cr Repository
}

func NewCreateService(r Repository) *CreateService {
	return &CreateService{Cr: r}
}

func (s *CreateService) Execute(d CreateDTO) (*entities.Customer, error) {
	c := entities.NewCustomer()
	c.FiscalID = d.FiscalID
	c.CorporateName = d.CorporateName
	c.TradeName = d.TradeName
	c.Address = d.Address
	c.Number = d.Number
	c.City = d.City
	c.State = d.State
	c.Zip = d.Zip
	c.Complement = d.Complement

	isValid, err := c.IsValid()
	if !isValid {
		return nil, err
	}

	createdCustomer, err := s.Cr.Add(&c)
	if err != nil {
		return nil, err
	}

	return createdCustomer, nil
}
