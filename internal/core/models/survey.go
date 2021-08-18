package models

import "time"

type Survey struct {
	ID                   ID
	Customer             *Customer
	AnnualRevenue        float64
	QuantityIssuedNFe    int
	QuantityIssuedNFSe   int
	QuantityIssuedNFCe   int
	QuantityIssuedCTe    int
	NFeStatesCovered     []string
	NFSeCitiesCovered    []string
	MainProjectLanguage  string
	ExternalIntegrations []string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
