package postgres

import (
	"context"
	"fmt"
	"gazuberlandia"

	"github.com/jmoiron/sqlx"
)

type OrderService struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) gazuberlandia.OrderService {
	return &OrderService{db: db}
}

func (p *OrderService) CreateOrder(ctx context.Context, order *gazuberlandia.Order) error {
	tx := p.db.MustBegin()

	result := tx.MustExec("INSERT INTO orders (id_costumer, id_product, total_amount, type_payment) VALUES ($1, $2, $3, $4) returning *;",
		&order.Id_costumer, &order.Id_product, &order.Total_amount, &order.Type_payment)

	fmt.Println(result)

	return tx.Commit()

}
