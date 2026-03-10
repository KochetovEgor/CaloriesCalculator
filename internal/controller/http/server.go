package http

import (
	"CaloriesCalculator/internal/pkg/config"
	"CaloriesCalculator/internal/service"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
	"net/http"
)

type App struct {
	service *service.Service
}

func New(service *service.Service) *App {
	return &App{service: service}
}

func (a *App) Run(ctx context.Context, cfg config.Server) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.Test)

	server := &http.Server{
		Handler:      mux,
		Addr:         cfg.Address,
		ReadTimeout:  cfg.Timeout.Duration,
		WriteTimeout: cfg.Timeout.Duration,
		IdleTimeout:  cfg.IdleTimeout.Duration,
	}

	logger := mylog.FromContext(ctx)
	logger.Info("server succesfully started", "addr", server.Addr)

	err := server.ListenAndServe()
	return err
}

func (a *App) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
