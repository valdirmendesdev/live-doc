package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

func (a *App) CustomerRoutes() {
	g := a.router.Group("/customers")
	g.Get("/:id", a.findById)
}

func (a *App) findById(c *fiber.Ctx) error {
	id, err := types.ParseID(c.Params("id"))
	if err != nil {
		return err
	}

	repo, _ := postgres.NewCustomerRepository(a.db)
	customer, err := repo.GetById(id)

	return c.JSON(customer)
}
