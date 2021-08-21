package postgres

import (
	"context"
	"gazuberlandia"

	"github.com/jmoiron/sqlx"
)

type CustomerService struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) gazuberlandia.CostumerService {
	return &CustomerService{db: db}
}

func (p *CustomerService) CreateCustomer(ctx context.Context, customer *gazuberlandia.Costumer) error {
	tx := p.db.MustBegin()

	tx.MustExec("INSERT INTO customer (name, phone, cpf) VALUES($1, $2, $3)",
		&customer.Name, &customer.Phone, &customer.Cpf)

	tx.Commit()
	return nil
}

func (p *CustomerService) FindCostumerById(ctx context.Context, id int) ([]gazuberlandia.Costumer, error) {
	var customer []gazuberlandia.Costumer

	err := p.db.Select(&customer, "SELECT * from customer c where c.id = $1", id)

	if err != nil {
		return nil, err
	}
	return customer, nil
}
