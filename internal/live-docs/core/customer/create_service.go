package customer

import (
	"log"

	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
)

type CreateService struct {
	Cr Repository
}

func NewCreateService(r Repository) *CreateService {
	return &CreateService{Cr: r}
}

func (s *CreateService) Execute(d dto.CustomerCreate) (*entities.Customer, error) {
	c := dto.CustomerCreateDtoToEntity(d)
	log.Printf("%v", c)

	isValid, err := c.IsValid()
	if !isValid {
		return nil, err
	}

	createdCustomer, err := s.Cr.Add(&c)
	if err != nil {
		return nil, err
	}

	return createdCustomer, nil
}
