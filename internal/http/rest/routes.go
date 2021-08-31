package rest

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

func (a *App) Setup() {
	a.router.Use(logger.New())
	a.HealthRoutes()
	a.CustomerRoutes()
}

func (a *App) Serve() error {
	return a.router.Listen(":8080")
}