package models_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"testing"
)

func Test_NewCustomer(t *testing.T) {
	u := models.NewCustomer()
	require.NotNil(t, u)
	require.NotNil(t, u.ID)
	require.False(t, u.CreatedAt.IsZero())
	require.True(t, u.UpdatedAt.IsZero())
}

func Test_CustomerIsValid(t *testing.T) {
	u := models.NewCustomer()
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
