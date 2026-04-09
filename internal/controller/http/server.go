package http

import (
	"CaloriesCalculator/internal/pkg/config"
	"CaloriesCalculator/internal/service"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	service *service.Service
}

func New(service *service.Service) *App {
	return &App{service: service}
}

func (a *App) Run(ctx context.Context, cfg config.Server) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.NotExists)

	// users
	mux.HandleFunc("POST /login", a.Login)
	mux.HandleFunc("POST /register", a.Register)

	// products
	mux.HandleFunc("GET /products", bearerAuthMiddleware(a.ProductsGet))
	mux.HandleFunc("POST /products", bearerAuthMiddleware(a.ProductsPost))
	mux.HandleFunc("PUT /products", bearerAuthMiddleware(a.ProductsPut))
	mux.HandleFunc("DELETE /products", bearerAuthMiddleware(a.ProductsDelete))

	// rations
	mux.HandleFunc("GET /rations", bearerAuthMiddleware(a.RationsGet))
	mux.HandleFunc("POST /rations", bearerAuthMiddleware(a.RationsPost))
	mux.HandleFunc("PUT /rations", bearerAuthMiddleware(a.RationsPut))
	mux.HandleFunc("DELETE /rations", bearerAuthMiddleware(a.RationsDelete))

	// rations/products
	mux.HandleFunc("PATCH /rations/products", bearerAuthMiddleware(a.RationsProductsPatch))

	handler := logMiddleware(CORSMiddleware(mux))

	server := &http.Server{
		Handler:      handler,
		Addr:         ":8000",
		ReadTimeout:  cfg.Timeout.Duration,
		WriteTimeout: cfg.Timeout.Duration,
		IdleTimeout:  cfg.IdleTimeout.Duration,
	}

	logger := mylog.FromContext(ctx).With("addr", server.Addr)

	chErr := make(chan error)
	go func() {
		<-ctx.Done()
		logger.Info("shutting down server")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		chErr <- server.Shutdown(ctxTimeout)
	}()

	logger.Info("server started")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error starting server: %w", err)
	}

	if err := <-chErr; err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}

func (a *App) NotExists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "endpoint not exists")
}
