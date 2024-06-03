package main

import (
	"context"

	"rutube-task/internal/app"
	"rutube-task/internal/config"
	"rutube-task/internal/handler"
	"rutube-task/internal/repository"
	"rutube-task/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("LoadConfig failed: %v", err)
	}

	db, err := repository.InitDataBase(ctx,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPass,
	)
	if err != nil {
		logrus.Fatalf("InitDataBase failed: %v", err)
	}

	router := chi.NewRouter()

	rep := repository.NewRepository(db)

	service := service.NewService(cfg, rep)

	handler := handler.NewHandler(service)

	handler.Register(router)

	srv := new(app.Server)

	if err := srv.Run(cfg, router); err != nil {
		logrus.Fatalf("Run failed: %v", err)
	}

}
