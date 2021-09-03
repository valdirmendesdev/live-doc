package controllers_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
)

func Test_HealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := fiber.New()
	hc := controllers.NewHealth()
	app.Get("/health", hc.Status)

	res, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	assert.Equal(t, `{"message":"OK!"}`, buf.String())

}
