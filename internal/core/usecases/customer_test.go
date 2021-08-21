package usecases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"github.com/valdirmendesdev/live-doc/internal/core/repositories/mocks"
	"github.com/valdirmendesdev/live-doc/internal/core/usecases"
)

func createCustomerUseCaseWithMock(t *testing.T) (*gomock.Controller, *mocks.MockCustomer, *usecases.Customer) {
	ctrl := gomock.NewController(t)
	repoMock := mocks.NewMockCustomer(ctrl)
	useCase := usecases.NewCustomer(repoMock)
	return ctrl, repoMock, useCase
}

func Test_CustomerUseCaseInstance(t *testing.T) {
	_, _, usecase := createCustomerUseCaseWithMock(t)

	require.NotNil(t, usecase)
	require.NotNil(t, usecase.Repo)
}

func Test_SuccessCustomerCreate(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(customer *models.Customer) (*models.Customer, error) {
			return customer, nil
		}).
		Times(1)

	dto := usecases.CustomerCreateRequest{
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

	customer, err := usecase.Create(dto)
	require.NotNil(t, customer)
	require.Nil(t, err)
	require.Equal(t, dto.FiscalID, customer.FiscalID)
	require.Equal(t, dto.CorporateName, customer.CorporateName)
	require.Equal(t, dto.TradeName, customer.TradeName)
	require.Equal(t, dto.Address, customer.Address)
	require.Equal(t, dto.Number, customer.Number)
	require.Equal(t, dto.City, customer.City)
	require.Equal(t, dto.State, customer.State)
	require.Equal(t, dto.Zip, customer.Zip)
	require.Equal(t, dto.Complement, customer.Complement)
}

func Test_RepositoryErrorCreateCustomer(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(customer *models.Customer) (*models.Customer, error) {
			return nil, errors.New("repository's error")
		}).
		Times(1)

	dto := usecases.CustomerCreateRequest{
		FiscalID:      "test",
		CorporateName: "test",
	}
	customer, err := usecase.Create(dto)
	require.Nil(t, customer)
	require.NotNil(t, err)
}

func Test_InvalidCustomerCreateCustomer(t *testing.T) {
	_, _, usecase := createCustomerUseCaseWithMock(t)

	dto := usecases.CustomerCreateRequest{}
	customer, err := usecase.Create(dto)
	require.Nil(t, customer)
	require.NotNil(t, err)
}

func Test_ListAllCustomers(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		ListAll(gomock.Any(), gomock.Any()).
		Return([]models.Customer{}, nil).
		Times(1)

	limit, page := 1,1
	customersList, err := usecase.List(limit, page)
	require.IsType(t, customersList, []models.Customer{})
	require.NotNil(t, customersList)
	require.Nil(t, err)

	repoMock.
		EXPECT().
		ListAll(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("repository's error")).
		Times(1)
	customersList, err = usecase.List(limit, page)
	require.Nil(t, customersList)
	require.NotNil(t, err)
	require.Error(t, err, "repository's error")
}