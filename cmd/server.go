package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valdirmendesdev/live-doc/internal/core/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "ok",
		})
	})

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}

	database.AutoMigrate(&db.Customer{})
	database.Create(&db.Customer{
		FiscalID: "25215772000171",
	})

	app.Listen(":8080")
}
