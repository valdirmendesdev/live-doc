package models

import (
	"errors"
	"time"
)

type Customer struct {
	ID            ID
	FiscalID      string    //TODO: Change the type to CNPJ
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

func NewCustomer(fiscalId string, corporateName string) *Customer {
	return &Customer{
		ID:        NewUUID(),
		FiscalID: fiscalId,
		CorporateName: corporateName,
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
