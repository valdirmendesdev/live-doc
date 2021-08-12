package customer_test

import (
	"github.com/stretchr/testify/assert"
	c "github.com/valdirmendesdev/live-doc/infrastructure/repository/customer"
	"github.com/valdirmendesdev/live-doc/usecases/customer"

	"testing"
)

func Test_CreatingService(t *testing.T) {
	repo := c.NewInMemory()
	service := customer.NewService(repo)
	assert.NotNil(t, service)
}
