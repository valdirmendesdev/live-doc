package entities

import (
	"errors"
	"time"

	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Survey struct {
	Model
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

func (s *Survey) IsValid() (bool, error) {

	if s.MainProjectLanguage == "" {
		return false, errors.New("Main project language must be filled")
	}
	return true, nil
}

func (s *Survey) TableName() string {
	return "surveys"
}
