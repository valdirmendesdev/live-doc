package postgres

import (
	"database/sql"
	g "github.com/jeffotoni/gconcat"

	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/core/models"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) (CustomerRepository, error) {
	return CustomerRepository{
		db: db,
	}, nil
}

func (r *CustomerRepository) Create(c *models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (r *CustomerRepository) FindById(id models.ID) (*models.Customer, error) {
	c := new(models.Customer)

	sqlStatement := g.Concat(
		"SELECT id, fiscal_id, corporate_name, ",
		"trade_name, address, number, city, state, ",
		"zip, complement, created_at, updated_at ",
		"FROM customers WHERE id=$1",
	)
	err := r.db.QueryRow(sqlStatement, id.String()).
	Scan(&c.ID, &c.FiscalID, &c.CorporateName, &c.TradeName, &c.Address,
		&c.Number, &c.City, &c.State, &c.Zip, &c.Complement, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CustomerRepository) ListAll(limit int, page int) ([]models.Customer, error) {
	panic("implement me")
}

func (r *CustomerRepository) Close() error {
	return r.db.Close()
}