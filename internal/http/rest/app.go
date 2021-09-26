package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/customer"
)

type App struct {
	Router *fiber.App
}

func NewApp() App {
	a := App{
		Router: fiber.New(),
	}
	a.Router.Use(logger.New())
	a.Router.Use(helmet.New())
	return a
}

func (a *App) HealthRoutes() {
	g := a.Router.Group("/health")
	c := controllers.NewHealth()
	g.Get("/", c.Status)
}

func (a *App) CustomersRoutes(r customer.Repository) {
	c := controllers.NewCustomer(r)
	g := a.Router.Group("/customers")
	g.Get("/", c.ListAll)
	g.Get("/:id", c.FindById)
	g.Post("/", c.Create)
}

func (a *App) Serve() error {
	return a.Router.Listen(":8080")
}
