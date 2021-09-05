package handler

import (
	"gazuberlandia"
	"gazuberlandia/api/handler/middlewares"
	"gazuberlandia/postgres"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	userService     gazuberlandia.UserService
	customerService gazuberlandia.CostumerService
	orderService    gazuberlandia.OrderService
}

func NewHandler(db *sqlx.DB) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Timeout(time.Second * 5))
	mux.Use(middleware.RequestID)
	mux.Use(middlewares.ApplicationJson)
	mux.Use(middleware.Logger)

	userRepository := postgres.NewUserRepository(db)
	customerRepository := postgres.NewCustomerRepository(db)
	orderRepository := postgres.NewOrderRepository(db)

	h := &Handler{
		userService:     userRepository,
		customerService: customerRepository,
		orderService:    orderRepository,
	}

	mux.Group(func(r chi.Router) {
		r.Post("/login", h.HandleLogin)
	})

	mux.Route("/customers", func(r chi.Router) {
		r.Post("/", h.HandleCreateCustomer)
		r.Get("/{customerId}", h.HandleFindCustomerById)
	})

	mux.Route("/orders", func(r chi.Router) {
		r.Post("/", h.HandlerCreateOrder)
	})

	return mux
}
