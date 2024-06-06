package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"rutube-task/internal/entity"
)

func (h *Handler) signUp(rw http.ResponseWriter, r *http.Request) {
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.AuthorizationService.ValidateLogin(r.Context(), input)
	if errors.Is(err, entity.ErrNotUniqueLogin) {
		http.Error(rw, "not unique login or empty login", http.StatusConflict)
		return
	}

	id, err := h.service.AuthorizationService.CreateUser(r.Context(), input)
	if err != nil {
		http.Error(rw, "err to create user", http.StatusInternalServerError)
		return
	}

	token, err := h.service.AuthorizationService.GenerateJWTToken(id)
	if err != nil {
		http.Error(rw, "err to generate token", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Authorization", "Bearer "+token)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(token))
}

func (h *Handler) singIn(rw http.ResponseWriter, r *http.Request) {
	var input entity.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.AuthorizationService.CheckData(r.Context(), input)
	if err != nil {
		http.Error(rw, "invalid login or password", http.StatusUnauthorized)
		return
	}

	token, err := h.service.AuthorizationService.GenerateJWTToken(id)
	if err != nil {
		http.Error(rw, "err to generate token", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Authorization", "Bearer "+token)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(token))
}
