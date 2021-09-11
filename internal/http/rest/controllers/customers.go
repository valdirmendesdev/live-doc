package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/customer"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type Customer struct {
	repo customer.Repository
}

func NewCustomer(r customer.Repository) Customer {
	return Customer{
		repo: r,
	}
}

func (cc *Customer) FindById(c *fiber.Ctx) error {
	id, err := types.ParseID(c.Params("id"))
	if err != nil {
		return err
	}

	s := customer.NewGetByIDService(cc.repo)
	customer, err := s.Execute(id)

	if err != nil {
		return err
	}

	return c.JSON(dto.CustomerToDto(*customer))
}

func (cc *Customer) ListAll(c *fiber.Ctx) error {

	s := customer.NewListAllService(cc.repo)
	list, err := s.Execute(0,0)
	if err != nil {
		return err
	}
	var result []dto.Customer

	for _, c := range list {
		result = append(result, dto.CustomerToDto(c))
	}

	return c.JSON(result)
}