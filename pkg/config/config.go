package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Timeout time.Duration `env:"SERVER_TIMEOUT"`
		Address string        `env:"SERVER_ADDRESS"`
		Port    uint16        `env:"SERVER_PORT"`
	}
}

func Load(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
