package db

import (
	"errors"
	"time"
)

type Customer struct {
	ID            ID         `gorm:"column:id;type:uuid;primaryKey"`
	FiscalID      string     `gorm:"column:fiscal_id"` //TODO: Change the type to CNPJ
	CorporateName string     `gorm:"column:corporate_name"`
	TradeName     string     `gorm:"column:trade_name"`
	Address       string     `gorm:"column:address"`
	Number        string     `gorm:"column:number"`
	City          string     `gorm:"column:city"`
	State         string     `gorm:"column:state"`
	Zip           string     `gorm:"column:zip"`
	Complement    string     `gorm:"column:complement"`
	CreatedAt     *time.Time `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
}

func NewCustomer() *Customer {
	return &Customer{
		ID: NewUUID(),
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