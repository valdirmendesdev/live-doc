package postgres_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/config"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/storage/postgres"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock(t *testing.T) (sqlmock.Sqlmock, postgres.CustomerRepository, error) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	config.DB, err = gorm.Open(pg.New(pg.Config{
		Conn: dbConn,
	}))

	if err != nil {
		return nil, postgres.CustomerRepository{}, err
	}

	repo := postgres.NewCustomerRepository()

	return mock, repo, err
}

func Test_FindCustomerById(t *testing.T) {
	mock, repo, err := NewMock(t)
	if err != nil {
		t.Log(err)
		return
	}

	query := `SELECT * FROM "customers" WHERE id = $1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT 1`

	c := entities.NewCustomer()

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

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(c.ID.String()).
		WillReturnRows(rows)

	customer, err := repo.GetById(c.ID)
	require.NotNil(t, customer)
	require.Nil(t, err)
}

func Test_GetAll(t *testing.T) {
	mock, repo, err := NewMock(t)
	if err != nil {
		t.Log(err)
		return
	}

	query := `SELECT * FROM "customers" WHERE "customers"."deleted_at" IS NULL`

	c := entities.NewCustomer()

	mock.
		ExpectQuery(regexp.QuoteMeta(query)). //Ignore query
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

	customers, err := repo.All(0, 0)
	require.NotNil(t, customers)
	require.Nil(t, err)
}
