package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server Server `json:"http_server"`
	// TODO: storage
}

type Server struct {
	Address string `json:"address"`
	// TODO: timeouts
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file %s: %w", configPath, err)
	}
	defer file.Close()

	cfg := &Config{}
	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("error decoding config file %s: %w", configPath, err)
	}

	return cfg, nil
}
