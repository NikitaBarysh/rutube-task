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

func (h *Handler) Register(router *chi.Mux) {
	router.Post("/register", h.signUp)
	router.Post("/login", h.singIn)
	router.Post("/set-employees", h.setEmployees)
	router.Get("/get-employees", h.getEmployees)

	router.Route("/", func(router chi.Router) {
		router.Use(h.AuthorizationMiddleware)
		router.Post("/subscribe", h.subscribe)
		router.Post("/unsubscribe", h.unsubscribe)
	})
}
