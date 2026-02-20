package main

import (
	"CaloriesCalculator/internal/config"
	"CaloriesCalculator/internal/mylog"
	"fmt"
	"log"
	"log/slog"
	"os"
)

const congigBackendPath = "config/backend.json"

func main() {
	mylog.InitLogger(os.Stdout)
	slog.Info("logger initialized")

	cfg, err := config.LoadConfig(congigBackendPath)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("configs loaded")

	fmt.Println(cfg.Server.Address)
}
