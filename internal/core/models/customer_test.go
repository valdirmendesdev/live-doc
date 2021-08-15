package models_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/core/models"
	"testing"
)

func createNewCustomer(fiscalID string) *models.Customer {
	user := models.NewCustomer()
	user.FiscalID = fiscalID
	return user
}

func Test_NewCustomer(t *testing.T) {
	u := createNewCustomer("")
	assert.NotNil(t, u)
	assert.NotNil(t, u.ID)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_IsValid(t *testing.T) {
	u := createNewCustomer("")

	isValid, err := u.IsValid()
	assert.False(t, isValid)
	assert.EqualError(t, err, "fiscal ID must be filled")

	u.FiscalID = "39055037000152"
	isValid, err = u.IsValid()
	assert.False(t, isValid)
	assert.EqualError(t, err, "corporate name must be filled")

}