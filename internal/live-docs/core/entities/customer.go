package entities

import (
	"errors"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"time"
)

type Customer struct {
	ID            types.ID
	FiscalID      string //TODO: Change the type to CNPJ
	CorporateName string
	TradeName     string
	Address       string
	Number        string
	City          string
	State         string
	Zip           string
	Complement    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewCustomer() Customer {
	return Customer{
		ID:        types.NewID(),
		CreatedAt: time.Now(),
	}
}

func (c *Customer) IsValid() (bool, error) {
	if c.FiscalID == "" {
		return false, errors.New("fiscal ID must be filled")
	}
	if c.CorporateName == "" {
		return false, errors.New("corporate name must be filled")
	}
	return true, nil
}