package handler

import (
	"rutube-task/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func NewHandler(newService *service.Service) *Handler {
	return &Handler{
		service: newService,
	}
}

func (h *Handler) Register(router *chi.Mux) {}
