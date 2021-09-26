package controllers_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/http/rest"
)

func Test_HealthCheck(t *testing.T) {
	app := rest.NewApp()
	app.HealthRoutes()

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Router.Test(req)
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
