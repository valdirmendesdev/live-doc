package entities

import (
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"time"
)

type Survey struct {
	ID                   types.ID
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
	IntegrationSystems   []*IntegrationSystem
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
