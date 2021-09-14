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

/* HANDLER -> qualquer estrutura que embedding http.Handler ou implementa os metodos de http.Handler
 */

/*
	ROUTER: -> interessante é usar uma unica instancia injetar essa instancia em funções
	que fazem o registro dos paths. porque dessa forma se tiver muitos dominos
	como: user, customer, auth etc. cada um deles terá uma função(constructor,factory)
	que recebe a instancia do router e registra os paths especificos.

	essa instancia entao deve ser criada em um função q tenha um nivel mais alto,
	como a main ou alguma função que tenha visão e consiga chamar todas que dependem da
	injeção do router.

	esse router tem que estar contido dentro de uma estrutura, como Server ou Application.
*/

/*
	SERVICES/BUSINESS-LOGIC: -> precisa da injeção de dependencias como db, logger.
	1º Criar uma struct Global, como Application ou Server. fazer o emmbed dos serviços
	nessa struct, que precisaram ser instanciados e injetado as instancias de db, logger.
*/

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
