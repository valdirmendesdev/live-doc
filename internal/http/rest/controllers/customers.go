package controllers

import (
	"net/http"

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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	s := customer.NewGetByIDService(cc.repo)
	customer, err := s.Execute(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(dto.EntityToCustomerViewDto(*customer))
}

func (cc *Customer) ListAll(c *fiber.Ctx) error {

	s := customer.NewListAllService(cc.repo)
	list, err := s.Execute(0, 0)
	if err != nil {
		return err
	}
	var result []dto.CustomerView

	for _, c := range list {
		result = append(result, dto.EntityToCustomerViewDto(c))
	}

	return c.JSON(result)
}

func (cc *Customer) Create(c *fiber.Ctx) error {

	var cDTO dto.CustomerCreate
	err := c.BodyParser(&cDTO)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	s := customer.NewCreateService(cc.repo)
	newCustomer, _ := s.Execute(cDTO)

	return c.
		Status(http.StatusCreated).
		JSON(dto.EntityToCustomerViewDto(*newCustomer))
}
