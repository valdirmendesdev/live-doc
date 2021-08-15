package db_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/core/db"
	"testing"
)

func createNewCustomer(fiscalID string) *db.Customer {
	user := db.NewCustomer()
	user.FiscalID = fiscalID
	return user
}

func Test_NewCustomer(t *testing.T) {
	u := createNewCustomer("")
	assert.NotNil(t, u)
	assert.NotNil(t, u.ID)
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