package customer

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type GetByIDService struct {
	Cr Repository
}

func NewGetByIDService(r Repository) *GetByIDService {
	return &GetByIDService{Cr: r}
}

func (s *GetByIDService) Execute(id types.ID) (*entities.Customer, error) {
	c, err := s.Cr.GetById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}