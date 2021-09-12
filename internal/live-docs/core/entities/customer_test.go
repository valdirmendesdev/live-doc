package entities_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"testing"
)

func Test_TableName(t *testing.T) {
	customer := entities.NewCustomer()
	require.Equal(t, "customers", customer.TableName())
}

func Test_NewCustomer(t *testing.T) {
	u := entities.NewCustomer()
	require.NotNil(t, u)
	require.NotNil(t, u.ID)
	require.False(t, u.CreatedAt.IsZero())
	require.True(t, u.UpdatedAt.IsZero())
}

func Test_CustomerIsValid(t *testing.T) {
	u := entities.NewCustomer()
	isValid, err := u.IsValid()
	require.False(t, isValid)
	require.NotNil(t, err)
	require.EqualError(t, err, "fiscal ID must be filled")

	u.FiscalID = "39055037000152"
	isValid, err = u.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "corporate name must be filled")

	u.CorporateName = "Corporate Name"
	isValid, err = u.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)
}
