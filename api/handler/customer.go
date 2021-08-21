package handler

import (
	"encoding/json"
	"gazuberlandia"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer *gazuberlandia.Costumer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, "Error decode Json", http.StatusBadRequest)
		return
	}

	err = h.customerService.CreateCustomer(r.Context(), customer)

	if err != nil {
		http.Error(w, "Error Create Customer", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"ok": "Sucess"})

	if err != nil {
		log.Print("Error", err)
	}
}

func (h *Handler) HandleFindCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId, err := strconv.Atoi(chi.URLParam(r, "customerId"))

	if err != nil {
		http.Error(w, "Error get customerId", http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.FindCostumerById(r.Context(), customerId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if customer == nil {
		err = json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(&customer)

	if err != nil {
		log.Print("Error", err)
	}
}
