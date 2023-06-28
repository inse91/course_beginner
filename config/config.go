package config

import (
	"encoding/json"
	"os"
)

const (
	defaultPort = "8080"
)

type Config struct {
	Port string
}

func Get() (Config, error) {
	bts, err := os.ReadFile("config/config.json")
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err = json.Unmarshal(bts, &cfg); err != nil {
		return Config{}, err
	}

	if cfg.Port == "" {
		cfg.Port = defaultPort
	}

	return cfg, nil
}
