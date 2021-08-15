package handler

import (
	"encoding/json"
	"fmt"
	"gazuberlandia"
	"net/http"
	"strconv"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
)

func (s *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user *gazuberlandia.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error invalid to struct from json", http.StatusBadRequest)
		return
	}

	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	user.Password = hash

	err = s.userService.CreateUser(r.Context(), user)

	if err != nil {
		http.Error(w, "Error create User", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Handler) HandlerFindUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))

	if err != nil {
		fmt.Println("Error get user Id")
	}

	user, _ := s.userService.FindUserById(r.Context(), userId)

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		fmt.Println("Error converted value in json")
	}

}

func (s *Handler) HandlerFindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := s.userService.FindAllUsers(r.Context())

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		fmt.Println("Error converted value in json")
	}
}
