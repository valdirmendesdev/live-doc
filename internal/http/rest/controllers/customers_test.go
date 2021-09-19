package controllers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	dto2 "github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"gorm.io/gorm"

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

func newCustomerToTest() entities.Customer {
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
	return c
}

func checkEnpointHandler(t *testing.T, app *fiber.App, method, url string, expectedStatusCode int, expectedBody string) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	res, err := app.Test(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	assert.Equal(t, expectedStatusCode, res.StatusCode)
	assert.Equal(t, expectedBody, string(body))
	return nil
}

func toJsonString(i interface{}) string {
	j, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(j)
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
			name:       "CustomerView Not Found",
			id:         types.NewID().String(),
			error:      gorm.ErrRecordNotFound,
			statusCode: http.StatusNotFound,
			body:       `{"error":"record not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, cc := createCustomerController(t)

			if test.error != nil {
				id, err := types.ParseID(test.id)
				if err != nil {
					t.Log(err)
					return
				}
				r.EXPECT().
					GetById(id).
					Return(nil, test.error)
			}

			app := fiber.New()
			app.Get("/:id", cc.FindById)
			err := checkEnpointHandler(t, app, "GET", "/"+test.id, test.statusCode, test.body)
			if err != nil {
				t.Log(err)
				return
			}
		})
	}

	t.Run("Success way", func(t *testing.T) {
		r, cc := createCustomerController(t)
		c := newCustomerToTest()
		r.EXPECT().
			GetById(c.ID).
			Return(&c, nil)

		cDTO := dto2.EntityToCustomerViewDto(c)

		app := fiber.New()
		app.Get("/:id", cc.FindById)
		err := checkEnpointHandler(t, app, "GET", "/"+c.ID.String(), http.StatusOK, toJsonString(cDTO))
		if err != nil {
			t.Log(err)
			return
		}
	})
}

func Test_CustomerListAll(t *testing.T) {
	t.Run("All Customers", func(t *testing.T) {
		r, cc := createCustomerController(t)
		c := newCustomerToTest()
		customers := []entities.Customer{c}
		dtoResponse := []dto.CustomerView{dto.EntityToCustomerViewDto(c)}
		r.
			EXPECT().
			All(gomock.Any(), gomock.Any()).
			Return(customers, nil)

		app := fiber.New()
		app.Get("/", cc.ListAll)
		err := checkEnpointHandler(t, app, "GET", "/", http.StatusOK, toJsonString(dtoResponse))
		if err != nil {
			t.Log(err)
			return
		}
	})
}
