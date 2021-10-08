package customer_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	cs "github.com/valdirmendesdev/live-doc/internal/live-docs/core/customer"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/mocks"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

func createGetByIDService(t *testing.T) (*gomock.Controller, *mocks.MockCustomer, *cs.GetByIDService) {
	c := gomock.NewController(t)
	r := mocks.NewMockCustomer(c)
	s := cs.NewGetByIDService(r)
	return c, r, s
}

func Test_GetById(t *testing.T) {
	_, r, s := createGetByIDService(t)

	r.
		EXPECT().
		GetById(gomock.Any()).
		DoAndReturn(func(id types.ID) (*entities.Customer, error) {
			if id == uuid.Nil {
				return nil, errors.New(errorRepositoryText)
			}
			return &entities.Customer{}, nil
		}).
		Times(2)

	var id types.ID
	c, err := s.Execute(id)
	require.IsType(t, c, &entities.Customer{})
	require.Nil(t, c)
	require.NotNil(t, err)
	require.EqualError(t, err, errorRepositoryText)

	id = types.NewID()
	c, err = s.Execute(id)
	require.IsType(t, c, &entities.Customer{})
	require.NotNil(t, c)
	require.Nil(t, err)
}
