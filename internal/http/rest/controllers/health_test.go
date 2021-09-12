package controllers_test

import (
	"io"
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

	res, err := app.Test(req)
	if err != nil {
		t.Log(err)
		return
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Log(err)
		return
	}

	assert.Equal(t, `{"message":"OK!"}`, string(body))

}
