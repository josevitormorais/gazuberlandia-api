package handler

import (
	"context"
	"gazuberlandia"
	"gazuberlandia/postgres"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	gazuberlandia.CostumerService
	gazuberlandia.OrderService
	gazuberlandia.UserService
}

func NewApplication(ctx context.Context, conn *sqlx.DB) *Application {

	return &Application{
		UserService:     postgres.NewUserRepository(conn),
		OrderService:    postgres.NewOrderRepository(conn),
		CostumerService: postgres.NewCustomerRepository(conn),
	}
}
