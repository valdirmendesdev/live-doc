package controllers_test

import (
	"net/http"
	"testing"

	"github.com/valdirmendesdev/live-doc/internal/http/rest"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"gorm.io/gorm"

	"github.com/golang/mock/gomock"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

func createCustomerRepository(t *testing.T) *mocks.MockCustomer {
	c := gomock.NewController(t)
	r := mocks.NewMockCustomer(c)
	return r
}

func setupCustomerControllers(t *testing.T) (rest.App, *mocks.MockCustomer) {
	r := createCustomerRepository(t)
	a := rest.NewApp()
	a.CustomersRoutes(r)
	return a, r
}

func newCustomerToTest() entities.Customer {
	c := entities.NewCustomer()
	c.FiscalID = "fiscalID"
	c.CorporateName = "corporateName"
	c.TradeName = "tradeName"
	c.Address = "address"
	c.Number = "number"
	c.City = "city"
	c.State = "state"
	c.Zip = "zip"
	c.Complement = "complement"
	return c
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
			app, r := setupCustomerControllers(t)

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

			err := testPath(t, app.Router, http.MethodGet, "/customers/"+test.id, nil, test.statusCode, test.body)
			if err != nil {
				t.Log(err)
				return
			}
		})
	}

	t.Run("Success way", func(t *testing.T) {
		app, r := setupCustomerControllers(t)
		c := newCustomerToTest()
		r.EXPECT().
			GetById(c.ID).
			Return(&c, nil)

		cDTO := dto.EntityToCustomerViewDto(c)

		err := testPath(t, app.Router, http.MethodGet, "/customers/"+c.ID.String(), nil, http.StatusOK, toJsonString(cDTO))
		if err != nil {
			t.Log(err)
			return
		}
	})
}

func Test_CustomerListAll(t *testing.T) {
	t.Run("All Customers", func(t *testing.T) {
		app, r := setupCustomerControllers(t)
		c := newCustomerToTest()
		customers := []entities.Customer{c}
		dtoResponse := []dto.CustomerView{dto.EntityToCustomerViewDto(c)}
		r.
			EXPECT().
			All(gomock.Any(), gomock.Any()).
			Return(customers, nil)

		err := testPath(t, app.Router, http.MethodGet, "/customers/", nil, http.StatusOK, toJsonString(dtoResponse))
		if err != nil {
			t.Log(err)
			return
		}
	})
}

func Test_CustomerCreate(t *testing.T) {
	app, r := setupCustomerControllers(t)

	cDTO := dto.CustomerCreate{
		FiscalID:      "fiscalID",
		CorporateName: "corporateName",
		TradeName:     "tradeName",
		Address:       "address",
		Number:        "number",
		City:          "city",
		State:         "state",
		Zip:           "zip",
		Complement:    "complement",
	}

	c := newCustomerToTest()

	r.EXPECT().
		Add(gomock.Any()).
		Return(&c, nil)

	err := testPath(t, app.Router, http.MethodPost, "/customers", cDTO, http.StatusCreated, toJsonString(dto.EntityToCustomerViewDto(c)))
	if err != nil {
		t.Log(err)
		return
	}
}
