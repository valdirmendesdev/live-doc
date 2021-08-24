package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valdirmendesdev/live-doc/infra/db/postgres"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

func FindById(c *fiber.Ctx) error {
	id, err := types.ParseID(c.Params("id"))
	if err != nil {
		return err
	}

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	repo, _ := postgres.NewCustomerRepository(db)
	customer, err := repo.FindById(id)

	return c.JSON(customer)
}
