package main

import (
	"CaloriesCalculator/internal/controller/http"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/internal/pkg/config"
	"CaloriesCalculator/internal/service"
	"CaloriesCalculator/internal/storage/postgres"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"log/slog"
	"os"
)

const (
	configBackendPath = "config/backend.json"
	secretKeyPath     = "config/HS256key.txt"
)

func main() {
	mylog.InitLogger(os.Stdout, slog.LevelDebug)
	slog.Info("logger initialized")
	ctx := mylog.NewContext(context.Background(), slog.Default())

	cfg, err := config.LoadConfig(configBackendPath)
	if err != nil {
		slog.ErrorContext(mylog.ErrToContext(ctx, err), err.Error())
		os.Exit(1)
	}
	slog.Info("configs loaded")

	secretKey, err := config.LoadSecretKey(secretKeyPath)
	if err != nil {
		slog.ErrorContext(mylog.ErrToContext(ctx, err), err.Error())
		os.Exit(1)
	}
	auth.SetKey(secretKey)
	slog.Info("secret key settled")

	postgresPool, err := postgres.NewPool(ctx, cfg.Storage)

	if err != nil {
		slog.ErrorContext(mylog.ErrToContext(ctx, err), err.Error())
		os.Exit(1)
	}

	userStorage := postgres.NewUserStorage(postgresPool)
	productStorage := postgres.NewProductStorage(postgresPool)

	service := service.New(userStorage, productStorage)

	defer func() {
		service.Close()
		slog.Info("storage succesfully closed")
	}()

	if err := service.Init(ctx); err != nil {
		slog.ErrorContext(mylog.ErrToContext(ctx, err), err.Error())
		os.Exit(1)
	}
	slog.Info("service initialized")

	app := http.New(service)
	if err := app.Run(ctx, cfg.Server); err != nil {
		slog.ErrorContext(mylog.ErrToContext(ctx, err), err.Error())
	}

	//service.TestProduct(ctx)
}
