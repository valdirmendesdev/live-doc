package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/http/rest"
)

func main() {

	app := rest.NewApp()
	app.Setup()
	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
