package entities

import (
	"errors"
	"time"

	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Survey struct {
	Model
	Customer             *Customer
	AnnualRevenue        float64 `gorm:"column:annual_revenue"`
	QuantityIssuedNFe    int     `gorm:"column:quantity_issued_nfe"`
	QuantityIssuedNFSe   int     `gorm:"column:quantity_issued_nfse"`
	QuantityIssuedNFCe   int     `gorm:"column:quantity_issued_nfce"`
	QuantityIssuedCTe    int     `gorm:"column:quantity_issued_cte"`
	MainProjectLanguage  string  `gorm:"column:main_project_language"`
	NFeStatesCovered     []string
	NFSeCitiesCovered    []string
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
		return false, errors.New("main project language must be filled")
	}
	return true, nil
}

func (s *Survey) TableName() string {
	return "surveys"
}
