package customer

import (
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
)

type ListAllService struct {
	Cr Repository
}

func NewListAllService(r Repository) *ListAllService {
	return &ListAllService{Cr: r}
}

func (s *ListAllService) Execute(limit int, page int) ([]entities.Customer, error) {
	customersList, err := s.Cr.All(limit, page)
	if err != nil {
		return nil, err
	}
	return customersList, nil
}