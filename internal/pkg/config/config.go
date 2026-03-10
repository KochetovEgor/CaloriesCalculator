package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Server  Server  `json:"http_server"`
	Storage Storage `json:"storage"`
}

type Server struct {
	Address     string   `json:"address"`
	Timeout     Duration `json:"timeout"`
	IdleTimeout Duration `json:"idle_timeout"`
}

type Storage struct {
	Url      string `json:"database_url"`
	MaxConns int    `json:"pool_max_conns"`
	MinConns int    `json:"pool_min_conns"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	duration, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}

var defaultConfig = Config{
	Server: Server{
		Address:     "localhost:8000",
		Timeout:     Duration{20 * time.Second},
		IdleTimeout: Duration{60 * time.Second},
	},
	Storage: Storage{
		MaxConns: 100,
		MinConns: 0,
	},
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file %s: %w", configPath, err)
	}
	defer file.Close()

	cfgCopy := defaultConfig
	cfg := &cfgCopy
	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("error decoding config file %s: %w", configPath, err)
	}

	return cfg, nil
}
