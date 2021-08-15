package gazuberlandia

import (
	"context"
	"time"
)

type Costumer struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Cpf        string    `json:"cpf"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CostumerService interface {
	CreateCustomer(ctx context.Context, customer *Costumer) error
	FindCostumerById(ctx context.Context, id int) ([]Costumer, error)
}
