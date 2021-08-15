package postgres

import (
	"context"
	"gazuberlandia"

	"github.com/jmoiron/sqlx"
)

func NewCustomerRepository(db *sqlx.DB) gazuberlandia.CostumerService {
	return &postgres{db: db}
}

func (p *postgres) CreateCustomer(ctx context.Context, customer *gazuberlandia.Costumer) error {
	tx := p.db.MustBegin()

	tx.MustExec("INSERT INTO customer (name, phone, cpf) VALUES($1, $2, $3)",
		&customer.Name, &customer.Phone, &customer.Cpf)

	tx.Commit()
	return nil
}

func (p *postgres) FindCostumerById(ctx context.Context, id int) ([]gazuberlandia.Costumer, error) {
	var customer []gazuberlandia.Costumer

	err := p.db.Select(&customer, "SELECT * from customer c where c.id = $1", id)

	if err != nil {
		return nil, err
	}
	return customer, nil
}
