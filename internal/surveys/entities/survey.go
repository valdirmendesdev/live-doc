package entities

import (
	c "github.com/valdirmendesdev/live-doc/internal/customers/entities"
	"github.com/valdirmendesdev/live-doc/internal/integration_systems/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"time"
)

type Survey struct {
	ID                   types.ID
	Customer             *c.Customer
	AnnualRevenue        float64
	QuantityIssuedNFe    int
	QuantityIssuedNFSe   int
	QuantityIssuedNFCe   int
	QuantityIssuedCTe    int
	NFeStatesCovered     []string
	NFSeCitiesCovered    []string
	MainProjectLanguage  string
	ExternalIntegrations []string
	IntegrationSystems   []*entities.IntegrationSystem
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
