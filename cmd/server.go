package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/customers/handlers"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "ok",
		})
	})

	//database, err := gorm.Open(sqlite.Open("test.models"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to open database")
	//}
	//
	//database.AutoMigrate(&models.Customer{})
	//database.Create(&models.Customer{
	//	FiscalID: "25215772000171",
	//})

	app.Get("/customers/:id", handlers.FindById)

	app.Listen(":8080")

}
