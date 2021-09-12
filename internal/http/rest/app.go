package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valdirmendesdev/live-doc/internal/config"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
)

type App struct {
	router *fiber.App
}

func NewApp() App {
	return App{
		router: fiber.New(),
	}
}

func (a *App) healthRoutes() {
	g := a.router.Group("/health")
	c := controllers.NewHealth()
	g.Get("/", c.Status)
}

func (a *App) customersRoutes() error {
	// db, err := gorm.Open(pg.New(pg.Config{
	// 	Conn: a.db,
	// }))
	// if err != nil {
	// 	return err
	// }
	// db.AutoMigrate(&entities.Customer{})
	r := postgres.NewCustomerRepository()
	c := controllers.NewCustomer(&r)
	g := a.router.Group("/customers")
	g.Get("/", c.ListAll)
	g.Get("/:id", c.FindById)
	return nil
}

func (a *App) Setup() {
	config.InitDatabase()
	a.router.Use(logger.New())
	a.healthRoutes()
	a.customersRoutes()
}

func (a *App) Serve() error {
	return a.router.Listen(":8080")
}
