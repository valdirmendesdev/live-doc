package dto_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/http/rest/dto"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"testing"
)

func Test_CustomerToDto(t *testing.T) {
	c := entities.NewCustomer()
	c.FiscalID = "test"
	c.CorporateName = "test"
	c.TradeName = "test"
	c.Address = "test"
	c.Number = "test"
	c.City = "test"
	c.State = "test"
	c.Zip = "test"
	c.Complement = "test"

	d := dto.CustomerToDto(c)
	require.Equal(t, c.FiscalID, d.FiscalID)
	require.Equal(t, c.CorporateName, d.CorporateName)
	require.Equal(t, c.TradeName, d.TradeName)
	require.Equal(t, c.Address, d.Address)
	require.Equal(t, c.Number, d.Number)
	require.Equal(t, c.City, d.City)
	require.Equal(t, c.State, d.State)
	require.Equal(t, c.Zip, d.Zip)
	require.Equal(t, c.Complement, d.Complement)
}
