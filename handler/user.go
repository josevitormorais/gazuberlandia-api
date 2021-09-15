package handler

import (
	"encoding/json"
	"fmt"
	"gazuberlandia"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)


func (h *HttpServer) RegisterUserRoutes(r *chi.Mux) {
	r.Get("/user/{userId}", h.HandlerFindUserById)
}

func (h *HttpServer) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user *gazuberlandia.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		ResponseError(w, &gazuberlandia.AppError{
			Code:    gazuberlandia.UNPROCESSABLEENTITY,
			Message: "The format json is invalid",
			Err:     err,
		})
		return
	}

	err = h.userService.CreateUser(r.Context(), user)

	if err != nil {
		if ok := strings.Contains(err.Error(), "already exists"); !ok {
			ResponseError(w, &gazuberlandia.AppError{
				Code:    gazuberlandia.CONFLICT,
				Message: "Email already exists",
				Err:     err,
			})
		}
		ResponseError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HttpServer) HandlerFindUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))

	if err != nil {
		fmt.Println("Error get user Id")
	}

	user, _ := h.userService.FindUserById(r.Context(), userId)

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		fmt.Println("Error converted value in json")
	}

}

func (h *Handler) HandlerFindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := h.userService.FindAllUsers(r.Context())

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		fmt.Println("Error converted value in json")
	}
}
