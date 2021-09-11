package dto

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"time"
)

type Customer struct {
	ID            types.ID  `json:"id"`
	FiscalID      string    `json:"fiscalId"`
	CorporateName string    `json:"corporateName"`
	TradeName     string    `json:"tradeName"`
	Address       string    `json:"address"`
	Number        string    `json:"number"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	Complement    string    `json:"complement"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func CustomerToDto(c entities.Customer) Customer {
	return Customer{
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