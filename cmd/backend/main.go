package main

import (
	"CaloriesCalculator/internal/config"
	"fmt"
	"log"
)

const congigBackendPath = "config/backend.json"

func main() {
	fmt.Println("Hello world")
	cfg, err := config.LoadConfig(congigBackendPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.Server.Address)
}
