package controllers_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"
	"io"
	"net/http"
	"testing"
)

func createCustomerRepository(t *testing.T) *mocks.MockCustomer {
	c := gomock.NewController(t)
	r := mocks.NewMockCustomer(c)
	return r
}

func createCustomerController(t *testing.T) (*mocks.MockCustomer, controllers.Customer) {
	r := createCustomerRepository(t)
	c := controllers.NewCustomer(r)
	return r, c
}

func Test_CustomerFindByID_WithInvalidUUID(t *testing.T) {
	_, cc := createCustomerController(t)
	app := fiber.New()
	app.Get("/:id", cc.FindById)

	req, err := http.NewRequest("GET", "/abcd", nil)
	if err != nil {
		t.Log(err)
		return
	}

	res, err := app.Test(req)
	if err != nil {
		t.Log(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Log(err)
		return
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, string(body), `{"error":"invalid UUID length: 4"}`)
}
