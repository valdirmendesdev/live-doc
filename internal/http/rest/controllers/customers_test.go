package controllers_test

import (
	"encoding/json"
	dto2 "github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"gorm.io/gorm"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/controllers"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
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

func Test_CustomerFindByID_CustomerNotFound(t *testing.T) {
	r, cc := createCustomerController(t)
	app := fiber.New()
	app.Get("/:id", cc.FindById)

	id := types.NewID()
	r.EXPECT().
		GetById(id).
		Return(nil, gorm.ErrRecordNotFound)

	req, err := http.NewRequest("GET", "/"+id.String(), nil)
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

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, string(body), `{"error":"record not found"}`)
}

func Test_CustomerFindByID(t *testing.T) {
	r, cc := createCustomerController(t)
	app := fiber.New()
	app.Get("/:id", cc.FindById)

	c := entities.NewCustomer()
	c.FiscalID = "test"
	c.CorporateName = "test"
	c.TradeName = "test"
	c.Address = "test"
	c.Number = "test"
	c.City = "test"
	c.State = "test"
	c.Zip = "test"
	c.Complement = "test"
	r.EXPECT().
		GetById(c.ID).
		Return(&c, nil)

	cDTO := dto2.CustomerToDto(c)

	req, err := http.NewRequest("GET", "/"+c.ID.String(), nil)
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
	j, _ := json.Marshal(cDTO)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, string(body), string(j))
}
