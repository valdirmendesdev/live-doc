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

	d := dto.EntityToCustomerViewDto(c)
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

func Test_CreateCustomerDtoToCustomer(t *testing.T){
	d := dto.CustomerCreate{
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

	c := dto.CustomerCreateDtoToEntity(d)
	require.Equal(t, d.FiscalID, c.FiscalID)
	require.Equal(t, d.CorporateName, c.CorporateName)
	require.Equal(t, d.TradeName, c.TradeName)
	require.Equal(t, d.Address, c.Address)
	require.Equal(t, d.Number, c.Number)
	require.Equal(t, d.City, c.City)
	require.Equal(t, d.State, c.State)
	require.Equal(t, d.Zip, c.Zip)
	require.Equal(t, d.Complement, c.Complement)
}