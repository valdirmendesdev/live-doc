package customer_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	cs "github.com/valdirmendesdev/live-doc/internal/live-docs/core/customer"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"
)

func createListAllService(t *testing.T) (*gomock.Controller, *mocks.MockCustomer, *cs.ListAllService) {
	c := gomock.NewController(t)
	r := mocks.NewMockCustomer(c)
	s := cs.NewListAllService(r)
	return c, r, s
}

func Test_ListAllCustomers(t *testing.T) {
	_, r, s := createListAllService(t)

	r.
		EXPECT().
		All(gomock.Any(), gomock.Any()).
		DoAndReturn(func(limit int, page int) ([]entities.Customer, error) {
			if page == -1 {
				return nil, errors.New(errorRepositoryText)
			}
			return []entities.Customer{}, nil
		}).
		Times(2)

	limit, page := 10, 1
	list, err := s.Execute(limit, page)
	require.IsType(t, list, []entities.Customer{})
	require.NotNil(t, list)
	require.Nil(t, err)

	//Repository error
	page = -1
	list, err = s.Execute(limit, page)
	require.Nil(t, list)
	require.NotNil(t, err)
	require.EqualError(t, err, errorRepositoryText)
}
