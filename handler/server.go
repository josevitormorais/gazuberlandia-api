package handler

import (
	"encoding/json"
	"fmt"
	"gazuberlandia"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db     *sqlx.DB
	router *chi.Mux
	gazuberlandia.UserService
}

func NewServer() *Server {
	s := &Server{
		router: chi.NewMux(),
	}
	s.routes()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.router.Get("/usersAll", s.HandlerFindAllUsers)
	// s.router.Get("/users/{userId}", s.HandlerFindUserById)
}

func (s *Server) HandlerFindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := s.FindAllUsers(r.Context())

	err := json.NewEncoder(w).Encode(&users)

	if err != nil {
		fmt.Println("Error converted value in json")
	}
}
