package dto

import (
	"time"

	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type CustomerView struct {
	ID            types.ID  `json:"id"`
	FiscalID      string    `json:"fiscal_id"`
	CorporateName string    `json:"corporate_name"`
	TradeName     string    `json:"trade_name"`
	Address       string    `json:"address"`
	Number        string    `json:"number"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	Complement    string    `json:"complement"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type CustomerCreate struct {
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

func EntityToCustomerViewDto(c entities.Customer) CustomerView {
	return CustomerView{
		ID:            c.ID,
		FiscalID:      c.FiscalID,
		CorporateName: c.CorporateName,
		TradeName:     c.TradeName,
		Address:       c.Address,
		Number:        c.Number,
		City:          c.City,
		State:         c.State,
		Zip:           c.Zip,
		Complement:    c.Complement,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}

func CustomerCreateDtoToEntity(d CustomerCreate) entities.Customer {
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
	return c
}
