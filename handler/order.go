package handler

import (
	"encoding/json"
	"gazuberlandia"
	"log"
	"net/http"
)

func (h *Handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order *gazuberlandia.Order

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error decode JSON", http.StatusBadRequest)
		return
	}

	err = h.orderService.CreateOrder(r.Context(), order)

	if err != nil {
		http.Error(w, "Error to inserted order", http.StatusBadRequest)
		return
	}

}
