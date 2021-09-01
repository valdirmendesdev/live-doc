package controllers

import "github.com/gofiber/fiber/v2"

type Health struct {}

func NewHealth() Health {
	return Health{}
}

func (hc *Health) Status(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "OK!",
	})
}