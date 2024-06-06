package app

import (
	"context"
	"fmt"
	"net/http"

	"rutube-task/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg *config.Config, handler http.Handler) error {
	url := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	s.httpServer = &http.Server{
		Addr:    url,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
