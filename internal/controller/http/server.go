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

var publicPaths = map[string][]string{
	"/login": nil,
}

func (a *App) Run(ctx context.Context, cfg config.Server) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.Test)
	mux.HandleFunc("POST /login", a.Login)

	handler := logMiddleware(authMiddleware(mux))

	server := &http.Server{
		Handler:      handler,
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
	ctx := r.Context()
	logger := mylog.FromContext(ctx)
	user, ok := getUserFromContext(ctx)
	if !ok {
		respErrWithLog(w, "invalid user", http.StatusBadRequest, logger)
		return
	}
	fmt.Fprintf(w, "Hello world %s\n", user)
}
