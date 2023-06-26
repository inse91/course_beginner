package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Port string
}

func Get() (Config, error) {
	cfgFile, err := os.Open("config/config.json")
	if err != nil {
		return Config{}, err
	}

	bts, err := io.ReadAll(cfgFile)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err = json.Unmarshal(bts, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
