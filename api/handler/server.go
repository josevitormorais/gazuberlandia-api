package handler

import (
	"gazuberlandia"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	http.Handler
	userService gazuberlandia.UserService
}

func NewServer(options ...func(*HttpServer)) (*HttpServer, error) {
	s := HttpServer{}

	for _, optionsFn := range options {
		optionsFn(&s)
	}

	return &s, nil
}

func NewConfigUserHandler(service gazuberlandia.UserService) func(*HttpServer) {
	return func(srv *HttpServer) {
		router := chi.NewRouter()
		router.Get("/user/{userId}", srv.HandlerFindUserById)
		srv.Handler = router
		srv.userService = service
	}
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.router.ServeHTTP(w, r)
// }

// func (s *Server) routes() {
// 	s.router.Get("/usersAll", s.HandlerFindAllUsers)
// 	// s.router.Get("/users/{userId}", s.HandlerFindUserById)
// }

// func (s *Server) HandlerFindAllUsers(w http.ResponseWriter, r *http.Request) {
// 	users, _ := s.FindAllUsers(r.Context())

// 	err := json.NewEncoder(w).Encode(&users)

// 	if err != nil {
// 		fmt.Println("Error converted value in json")
// 	}
// }
