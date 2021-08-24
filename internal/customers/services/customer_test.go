package services_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/valdirmendesdev/live-doc/internal/customers/dto"
	"github.com/valdirmendesdev/live-doc/internal/customers/entities"
	"github.com/valdirmendesdev/live-doc/internal/customers/mocks"
	"github.com/valdirmendesdev/live-doc/internal/customers/services"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const errorReposityText = "repository's error"

func createCustomerUseCaseWithMock(t *testing.T) (*gomock.Controller, *mocks.MockCustomer, *services.Customer) {
	ctrl := gomock.NewController(t)
	repoMock := mocks.NewMockCustomer(ctrl)
	useCase := services.NewCustomer(repoMock)
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
		DoAndReturn(func(customer *entities.Customer) (*entities.Customer, error) {
			return customer, nil
		}).
		Times(1)

	createRequest := dto.CreateRequest{
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

	customer, err := usecase.Create(createRequest)
	require.NotNil(t, customer)
	require.Nil(t, err)
	require.Equal(t, createRequest.FiscalID, customer.FiscalID)
	require.Equal(t, createRequest.CorporateName, customer.CorporateName)
	require.Equal(t, createRequest.TradeName, customer.TradeName)
	require.Equal(t, createRequest.Address, customer.Address)
	require.Equal(t, createRequest.Number, customer.Number)
	require.Equal(t, createRequest.City, customer.City)
	require.Equal(t, createRequest.State, customer.State)
	require.Equal(t, createRequest.Zip, customer.Zip)
	require.Equal(t, createRequest.Complement, customer.Complement)
}

func Test_RepositoryErrorCreateCustomer(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(customer *entities.Customer) (*entities.Customer, error) {
			return nil, errors.New(errorReposityText)
		}).
		Times(1)

	createRequest := dto.CreateRequest{
		FiscalID:      "test",
		CorporateName: "test",
	}
	customer, err := usecase.Create(createRequest)
	require.Nil(t, customer)
	require.NotNil(t, err)
}

func Test_InvalidCustomerCreateCustomer(t *testing.T) {
	_, _, usecase := createCustomerUseCaseWithMock(t)

	createRequest := dto.CreateRequest{}
	customer, err := usecase.Create(createRequest)
	require.Nil(t, customer)
	require.NotNil(t, err)
}

func Test_ListAllCustomers(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		ListAll(gomock.Any(), gomock.Any()).
		DoAndReturn(func(limit int, page int) ([]entities.Customer, error) {
			if page == -1 {
				return nil, errors.New(errorReposityText)
			}
			return []entities.Customer{}, nil
		}).
		Times(2)

	limit, page := 10, 1
	customersList, err := usecase.ListAll(limit, page)
	require.IsType(t, customersList, []entities.Customer{})
	require.NotNil(t, customersList)
	require.Nil(t, err)

	//Repostiry error
	page = -1
	customersList, err = usecase.ListAll(limit, page)
	require.Nil(t, customersList)
	require.NotNil(t, err)
	require.EqualError(t, err, errorReposityText)
}

func Test_FindCustomerById(t *testing.T) {
	_, repoMock, usecase := createCustomerUseCaseWithMock(t)

	repoMock.
		EXPECT().
		FindById(gomock.Any()).
		DoAndReturn(func(id types.ID) (*entities.Customer, error) {
			if id == uuid.Nil {
				return nil, errors.New(errorReposityText)
			}
			return &entities.Customer{}, nil
		}).
		Times(2)

	var customerID types.ID
	customer, err := usecase.FindById(customerID)
	require.IsType(t, customer, &entities.Customer{})
	require.Nil(t, customer)
	require.NotNil(t, err)
	require.EqualError(t, err, errorReposityText)

	customerID = types.NewID()
	customer, err = usecase.FindById(customerID)
	require.IsType(t, customer, &entities.Customer{})
	require.NotNil(t, customer)
	require.Nil(t, err)
}
