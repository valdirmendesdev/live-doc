package rest

import "github.com/gofiber/fiber/v2"

func (a *App) HealthRoutes()  {
	g := a.router.Group("/health")
	g.Get("/", a.healthController)
}

func (a *App) healthController(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "ok",
	})
}