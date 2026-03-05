package main

import (
	"CaloriesCalculator/internal/config"
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/storage"
	"CaloriesCalculator/internal/storage/postgres"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
	"log/slog"
	"os"
)

const configBackendPath = "config/backend.json"

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

	PostgresStorage, err := postgres.New(ctx, cfg.Storage)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	storage.SetDefault(PostgresStorage)
	slog.Info("storage created")

	if err := storage.Init(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := storage.AddUser(ctx, domain.User{"Egor", "123"}); err != nil {
		slog.Error(err.Error())
	}

	if err := storage.DeleteUser(ctx, "egor"); err != nil {
		slog.Error(err.Error())
	}

	user, err := storage.SelectUser(ctx, "Egor")
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(user)

	defer func() {
		storage.Close()
		slog.Info("storage succesfully closed")
	}()
}
