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
	mux.HandleFunc("/", a.NotExists)

	// users
	mux.HandleFunc("POST /login", a.Login)
	mux.HandleFunc("POST /register", a.Register)

	//products
	mux.HandleFunc("GET /products", authMiddleware(a.ProductsGet))
	mux.HandleFunc("POST /products", authMiddleware(a.ProductsPost))
	mux.HandleFunc("PUT /products", authMiddleware(a.ProductsPut))
	mux.HandleFunc("DELETE /products", authMiddleware(a.ProductsDelete))

	//rations
	mux.HandleFunc("GET /rations", authMiddleware(a.RationsGet))
	mux.HandleFunc("POST /rations", authMiddleware(a.RationsPost))
	mux.HandleFunc("PUT /rations", authMiddleware(a.RationsPut))
	mux.HandleFunc("DELETE /rations", authMiddleware(a.RationsDelete))

	//rations/products
	mux.HandleFunc("PATCH /rations/products", authMiddleware(a.RationsProductsPatch))

	handler := logMiddleware(mux)

	server := &http.Server{
		Handler:      handler,
		Addr:         ":8000",
		ReadTimeout:  cfg.Timeout.Duration,
		WriteTimeout: cfg.Timeout.Duration,
		IdleTimeout:  cfg.IdleTimeout.Duration,
	}

	logger := mylog.FromContext(ctx)
	logger.Info("server succesfully started", "addr", server.Addr)

	err := server.ListenAndServe()
	return err
}

func (a *App) NotExists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "endpoint not exists")
}
