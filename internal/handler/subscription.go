package handler

import (
	"encoding/json"
	"net/http"

	"rutube-task/internal/entity"
)

func (h *Handler) subscribe(rw http.ResponseWriter, r *http.Request) {
	var input entity.SubscriptionRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(rw, "Err to read body", http.StatusBadRequest)
		return
	}

	currUser := r.Context().Value("user").(int)

	err = h.service.SubscriptionServiceInterface.Subscribe(r.Context(), currUser, input.Name)
	if err != nil {
		http.Error(rw, "Err to subscribe", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) unsubscribe(rw http.ResponseWriter, r *http.Request) {
	var input entity.SubscriptionRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(rw, "Err to read body", http.StatusBadRequest)
		return
	}

	currUser := r.Context().Value("user").(int)

	err = h.service.SubscriptionServiceInterface.Unsubscribe(r.Context(), currUser, input.Name)
	if err != nil {
		http.Error(rw, "Err to unsubscribe", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
