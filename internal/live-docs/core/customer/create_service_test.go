package customer_test

import (
	"errors"
	"testing"

	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	cs "github.com/valdirmendesdev/live-doc/internal/live-docs/core/customer"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const errorRepositoryText = "repository's error"

func createService(t *testing.T) (*gomock.Controller, *mocks.MockCustomer, *cs.CreateService) {
	c := gomock.NewController(t)
	r := mocks.NewMockCustomer(c)
	s := cs.NewCreateService(r)
	return c, r, s
}

func Test_CreateServiceInstance(t *testing.T) {
	_, _, s := createService(t)

	require.NotNil(t, s)
	require.NotNil(t, s.Cr)
}

func Test_SuccessCustomerCreate(t *testing.T) {
	_, r, s := createService(t)

	r.
		EXPECT().
		Add(gomock.Any()).
		DoAndReturn(func(customer *entities.Customer) (*entities.Customer, error) {
			return customer, nil
		}).
		Times(1)

	d := dto.CustomerCreate{
		FiscalID:      "test",
		CorporateName: "test",
		TradeName:     "test",
		Address:       "test",
		Number:        "test",
		City:          "test",
		State:         "test",
		Zip:           "test",
		Complement:    "test",
	}

	c, err := s.Execute(d)
	require.NotNil(t, c)
	require.Nil(t, err)
	require.Equal(t, d.FiscalID, c.FiscalID)
	require.Equal(t, d.CorporateName, c.CorporateName)
	require.Equal(t, d.TradeName, c.TradeName)
	require.Equal(t, d.Address, c.Address)
	require.Equal(t, d.Number, c.Number)
	require.Equal(t, d.City, c.City)
	require.Equal(t, d.State, c.State)
	require.Equal(t, d.Zip, c.Zip)
	require.Equal(t, d.Complement, c.Complement)
}

func Test_CreateRepositoryError(t *testing.T) {
	_, r, s := createService(t)

	r.
		EXPECT().
		Add(gomock.Any()).
		DoAndReturn(func(customer *entities.Customer) (*entities.Customer, error) {
			return nil, errors.New(errorRepositoryText)
		}).
		Times(1)

	d := dto.CustomerCreate{
		FiscalID:      "test",
		CorporateName: "test",
	}
	c, err := s.Execute(d)
	require.Nil(t, c)
	require.NotNil(t, err)
}

func Test_InvalidCreateInputData(t *testing.T) {
	_, _, s := createService(t)

	d := dto.CustomerCreate{}
	c, err := s.Execute(d)
	require.Nil(t, c)
	require.NotNil(t, err)
}
