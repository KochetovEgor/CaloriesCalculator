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
	mylog.InitLogger(os.Stdout)
	slog.Info("logger initialized")
	ctx := mylog.NewContext(context.Background(), slog.Default())

	cfg, err := config.LoadConfig(configBackendPath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("configs loaded")

	secretKey, err := config.LoadSecretKey(secretKeyPath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	auth.SetKey(secretKey)
	slog.Info("secret key settled")

	PostgresStorage, err := postgres.New(ctx, cfg.Storage)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	service := service.New(PostgresStorage)

	if err := service.Init(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	app := http.New(service)
	if err := app.Run(ctx, cfg.Server); err != nil {
		slog.Error(err.Error())
	}

	//service.Test(ctx)

	defer func() {
		service.Close()
		slog.Info("storage succesfully closed")
	}()
}
