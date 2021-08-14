package services_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/live-doc/internal/core/db"
	repo "github.com/valdirmendesdev/live-doc/internal/core/repositories/InMemory"
	"github.com/valdirmendesdev/live-doc/internal/core/services"
	"testing"
)

func newTestCustomer() *db.Customer {
	return &db.Customer{}
}

func Test_CreatingService(t *testing.T) {
	repo := repo.NewInMemory()
	service := services.NewCustomer(repo)
	assert.NotNil(t, service)
}