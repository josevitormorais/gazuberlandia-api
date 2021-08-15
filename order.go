package gazuberlandia

import (
	"context"
	"time"
)

type Order struct {
	ID           int       `json:"id"`
	Id_costumer  int       `json:"id_costumer"`
	Id_product   int       `json:"id_product"`
	Total_amount int       `json:"total_amount"`
	Type_payment string    `json:"type_payment"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type OrderService interface {
	CreateOrder(ctx context.Context, order *Order) error
}
