package entities

import (
	"time"

	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Survey struct {
	Model
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
}

func NewSurvey() Survey {
	return Survey{
		Model: Model{
			ID:        types.NewID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func (s *Survey) TableName() string {
	return "surveys"
}
