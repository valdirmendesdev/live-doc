package entities

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID            uuid.UUID `json:"id"`
	FiscalID      string    `json:"fiscal_id"`
	CorporateName string    `json:"corporate_name"`
	TradeName     string    `json:"trade_name"`
	Address       string    `json:"address"`
	Number        string    `json:"number"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	Complement    string    `json:"complement"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
