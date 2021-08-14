package db

import (
	"time"
)

type Customer struct {
	ID            ID         `gorm:"column:id;type:uuid;primaryKey"`
	FiscalID      string     `gorm:"column:fiscal_id"`
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
