package usecases_test

import (
	"errors"
	"github.com/google/uuid"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"github.com/valdirmendesdev/live-doc/internal/core/repositories/mocks"
	"github.com/valdirmendesdev/live-doc/internal/core/usecases"
)

const errorReposityText = "repository's error"

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
			return nil, errors.New(errorReposityText)
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
		DoAndReturn(func(limit int, page int) ([]models.Customer, error) {
			if page == -1 {
				return nil, errors.New(errorReposityText)
			}
			return []models.Customer{}, nil
		}).
		Times(2)

	limit, page := 10, 1
	customersList, err := usecase.ListAll(limit, page)
	require.IsType(t, customersList, []models.Customer{})
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
		DoAndReturn(func(id models.ID) (*models.Customer, error) {
			if id == uuid.Nil {
				return nil, errors.New(errorReposityText)
			}
			return &models.Customer{}, nil
		}).
		Times(2)

	var customerID models.ID
	customer, err := usecase.FindById(customerID)
	require.IsType(t, customer, &models.Customer{})
	require.Nil(t, customer)
	require.NotNil(t, err)
	require.EqualError(t, err, errorReposityText)

	customerID = models.NewUUID()
	customer, err = usecase.FindById(customerID)
	require.IsType(t, customer, &models.Customer{})
	require.NotNil(t, customer)
	require.Nil(t, err)

}
