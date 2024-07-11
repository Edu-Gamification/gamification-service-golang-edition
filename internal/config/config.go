package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	URI string `yaml:"uri"`
}

func ParseFromYaml() (*Config, error) {
	cfgPath := "config/config.yaml"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return nil, fmt.Errorf("load config from yaml: %w", err)
	}

	return &cfg, nil
}
