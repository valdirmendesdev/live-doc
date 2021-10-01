package entities

import (
	"errors"
	"time"

	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Customer struct {
	Model
	FiscalID      string `gorm:"column:fiscal_id"` //TODO: Change the type to CNPJ
	CorporateName string `gorm:"column:corporate_name"`
	TradeName     string `gorm:"column:trade_name"`
	Address       string `gorm:"column:address"`
	Number        string `gorm:"column:number"`
	City          string `gorm:"column:city"`
	State         string `gorm:"column:state"`
	Zip           string `gorm:"column:zip"`
	Complement    string `gorm:"column:complement"`
}

func NewCustomer() Customer {
	return Customer{
		Model: Model{
			ID:        types.NewID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
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

func (c *Customer) TableName() string {
	return "customers"
}
