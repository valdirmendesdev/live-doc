package models_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"testing"
)

func createNewCustomer() models.Customer {
	user := models.NewCustomer("123456789", "test")
	return user
}

func Test_NewCustomer(t *testing.T) {
	u := createNewCustomer()
	require.NotNil(t, u)
	require.NotNil(t, u.ID)
	require.NotNil(t, u.FiscalID)
	require.NotNil(t, u.CorporateName)
	require.False(t, u.CreatedAt.IsZero())
	require.True(t, u.UpdatedAt.IsZero())
}

func Test_CustomerIsValid(t *testing.T) {
	u := createNewCustomer()
	u.FiscalID = ""
	u.CorporateName = ""
	isValid, err := u.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "fiscal ID must be filled")

	u.FiscalID = "39055037000152"
	isValid, err = u.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "corporate name must be filled")
}
