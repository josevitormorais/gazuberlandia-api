package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {

	var login *loginRequest

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, "Error decode json", http.StatusBadRequest)
		return
	}
	time.Sleep(time.Second * 7)

	user, err := h.userService.FindUserByEmail(r.Context(), login.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	match, err := argon2id.ComparePasswordAndHash(login.Password, user.Password)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	isEqual := strings.Compare(user.Email, login.Email)

	if !match || isEqual != 0 {
		http.Error(w, "Error: email or password is invalid", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
