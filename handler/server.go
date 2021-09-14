package handler

import (
	"gazuberlandia"
	"gazuberlandia/postgres"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type HttpServer struct {
	http.Handler
	Version string
	server  *http.Server

	costumerService gazuberlandia.CostumerService
	orderService    gazuberlandia.OrderService
	userService     gazuberlandia.UserService
}

func NewServer(options ...func(*HttpServer)) *HttpServer {
	server := &HttpServer{}

	for _, optionsFn := range options {
		optionsFn(server)
	}

	return server
}

func NewRouter() func(*HttpServer) {
	return func(srv *HttpServer) {
		router := chi.NewRouter()
		router.Get("/version", srv.handleVersionApi)

		srv.RegisterUserRoutes(router)
		srv.RegisterOrdersRouter(router)
		srv.Handler = router
	}
}

func NewServices(db *sqlx.DB) func(*HttpServer) {
	return func(srv *HttpServer) {
		srv.userService = postgres.NewUserRepository(db)
		srv.orderService = postgres.NewOrderRepository(db)
		srv.costumerService = postgres.NewCustomerRepository(db)
	}
}

func (s *HttpServer) handleVersionApi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.Version = "0.0.1"
	w.Write([]byte(s.Version))
}
