package rest

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
)

type App struct {
	db *sql.DB
	router *fiber.App
}

func NewApp(d *sql.DB) App {
	return App{
		db: d,
		router: fiber.New(),
	}
}

func (a *App) healthRoutes()  {
	g := a.router.Group("/health")
	c := controllers.NewHealth()
	g.Get("/", c.Status)
}

func (a *App) customersRoutes() error {
	r, err := postgres.NewCustomerRepository(a.db)
	if err != nil {
		return nil
	}
	c := controllers.NewCustomer(&r)
	g := a.router.Group("/customers")
	g.Get("/", c.ListAll)
	g.Get("/:id", c.FindById)
	return nil
}

func (a *App) Setup() {
	a.router.Use(logger.New())
	a.healthRoutes()
	a.customersRoutes()
}

func (a *App) Serve() error {
	return a.router.Listen(":8080")
}