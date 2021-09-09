package postgres_test

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	g "github.com/jeffotoni/gconcat"
)

func NewMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, postgres.CustomerRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo, err := postgres.NewCustomerRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a customer repository", err)
	}

	return db, mock, repo
}

func newCustomer() *entities.Customer {
	return &entities.Customer{
		ID:            types.NewID(),
		FiscalID:      "",
		TradeName:     "",
		CorporateName: "",
		Address:       "",
		Number:        "",
		City:          "",
		State:         "",
		Zip:           "",
		Complement:    "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}

func Test_FindCustomerById(t *testing.T) {
	_, mock, repo := NewMock(t)

	query := g.Concat(
		"SELECT id, fiscal_id, corporate_name, ",
		"trade_name, address, number, city, state, ",
		"zip, complement, created_at, updated_at ",
		"FROM customers WHERE id=\\$1",
	)

	c := newCustomer()

	rows := sqlmock.NewRows([]string{"id",
		"fiscal_id", "corporate_name", "trade_name",
		"address", "number", "city", "state", "zip",
		"complement", "created_at", "updated_at"}).
		AddRow(
			c.ID,
			c.FiscalID,
			c.CorporateName,
			c.TradeName,
			c.Address,
			c.Number,
			c.City,
			c.State,
			c.Zip,
			c.Complement,
			c.CreatedAt,
			c.UpdatedAt,
		)

	mock.ExpectQuery(query).WithArgs(c.ID.String()).WillReturnRows(rows)

	customer, err := repo.GetById(c.ID)
	require.NotNil(t, customer)
	require.Nil(t, err)
}

func Test_GetAll(t *testing.T) {
	_, mock, repo := NewMock(t)

	query := g.Concat(
		"SELECT id, fiscal_id, corporate_name, ",
		"trade_name, address, number, city, state, ",
		"zip, complement, created_at, updated_at ",
		"FROM customers WHERE id IS NOT NULL",
	)

	c := newCustomer()

	mock.
		ExpectQuery(query).
		WillReturnRows(sqlmock.NewRows(
			[]string{"id",
				"fiscal_id", "corporate_name", "trade_name",
				"address", "number", "city", "state", "zip",
				"complement", "created_at", "updated_at",
			}).AddRow(
			c.ID,
			c.FiscalID,
			c.CorporateName,
			c.TradeName,
			c.Address,
			c.Number,
			c.City,
			c.State,
			c.Zip,
			c.Complement,
			c.CreatedAt,
			c.UpdatedAt),
		)

	customers, err := repo.All(0,0)
	require.NotNil(t, customers)
	require.Nil(t, err)
}
