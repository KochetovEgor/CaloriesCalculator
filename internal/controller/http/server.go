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
	mux.Handle("/", authMiddleware(a.Test))

	mux.HandleFunc("POST /login", a.Login)
	mux.HandleFunc("POST /register", a.Register)

	mux.Handle("POST /product/add", authMiddleware(a.ProductAdd))
	mux.Handle("DELETE /product/delete", authMiddleware(a.ProductDelete))
	mux.Handle("PUT /product/update", authMiddleware(a.ProductUpdate))
	mux.Handle("GET /product", authMiddleware(a.Product))

	mux.Handle("POST /ration/add", authMiddleware(a.RationAdd))
	mux.Handle("DELETE /ration/delete", authMiddleware(a.RationDelete))
	mux.Handle("PUT /ration/update", authMiddleware(a.RationUpdate))
	mux.Handle("GET /ration", authMiddleware(a.Ration))

	handler := logMiddleware(mux)

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
	user := getUserFromContext(ctx)
	fmt.Fprintf(w, "Hello world %v\n", user)
}
