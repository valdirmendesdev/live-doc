package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/config"
	"github.com/valdirmendesdev/live-doc/internal/http/rest"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
)

func main() {
	config.InitDatabase()

	cr := postgres.NewCustomerRepository()

	app := rest.NewApp()
	app.HealthRoutes()
	app.CustomersRoutes(&cr)

	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
