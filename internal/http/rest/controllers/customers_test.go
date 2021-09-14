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

func Test_CustomerFindByID(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		error      error
		statusCode int
		body       string
	}{
		{
			name:       "Invalid UUID",
			id:         "abcd",
			error:      nil,
			statusCode: http.StatusBadRequest,
			body:       `{"error":"invalid UUID length: 4"}`,
		},
		{
			name:       "Customer Not Found",
			id:         types.NewID().String(),
			error:      gorm.ErrRecordNotFound,
			statusCode: http.StatusNotFound,
			body:       `{"error":"record not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, cc := createCustomerController(t)
			app := fiber.New()
			app.Get("/:id", cc.FindById)

			if test.error != nil {
				id, err :=  types.ParseID(test.id)
				if err != nil {
					t.Log(err)
					return
				}
				r.EXPECT().
					GetById(id).
					Return(nil, test.error)
			}

			req, err := http.NewRequest("GET", "/"+test.id, nil)
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

			assert.Equal(t, test.statusCode, res.StatusCode)
			assert.Equal(t, string(body), test.body)
		})
	}
}

func Test_CustomerFindByID_Success(t *testing.T) {
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
