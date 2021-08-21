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
	mock := mocks.NewMockCustomer(ctrl)
	useCase := usecases.NewCustomer(mock)
	return ctrl, mock, useCase
}

func Test_CustomerUseCaseInstance(t *testing.T) {
	ctrl, _, useCase := createCustomerUseCaseWithMock(t)
	defer ctrl.Finish()

	require.NotNil(t, useCase)
	require.NotNil(t, useCase.Repo)
}

func Test_SuccessCustomerCreate(t *testing.T) {
	ctrl, mock, useCase := createCustomerUseCaseWithMock(t)
	defer ctrl.Finish()

	mock.
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

	customer, err := useCase.Create(dto)
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
	ctrl, mock, useCase := createCustomerUseCaseWithMock(t)
	defer ctrl.Finish()

	mock.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(customer *models.Customer) (*models.Customer, error) {
			return nil, errors.New("Error from repository")
		}).
		Times(1)

	dto := usecases.CustomerCreateRequest{
		FiscalID:      "test",
		CorporateName: "test",
	}
	customer, err := useCase.Create(dto)
	require.Nil(t, customer)
	require.NotNil(t, err)
}

func Test_InvalidCustomerCreateUseCase(t *testing.T) {
	ctrl, _, useCase := createCustomerUseCaseWithMock(t)
	defer ctrl.Finish()

	dto := usecases.CustomerCreateRequest{}
	customer, err := useCase.Create(dto)
	require.Nil(t, customer)
	require.NotNil(t, err)
}
