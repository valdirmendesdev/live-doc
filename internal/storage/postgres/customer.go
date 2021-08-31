package postgres

import (
	"database/sql"
	g "github.com/jeffotoni/gconcat"
	_ "github.com/lib/pq"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"github.com/valdirmendesdev/live-doc/internal/utils/types"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) (CustomerRepository, error) {
	return CustomerRepository{
		db: db,
	}, nil
}

func (r *CustomerRepository) Add(c *entities.Customer) (*entities.Customer, error) {
	panic("implement me")
}

func (r *CustomerRepository) GetById(id types.ID) (*entities.Customer, error) {
	c := new(entities.Customer)

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

func (r *CustomerRepository) All(limit int, page int) ([]entities.Customer, error) {
	panic("implement me")
}

func (r *CustomerRepository) Close() error {
	return r.db.Close()
}