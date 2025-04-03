package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port     int    `env:"PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"0.0.0.0"`
	Greeting string `env:"GREETING" env-default:"Hello"`
}

func NewConfig() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("не удалось инициализировать конфиг: %w", err)
	}

	return cfg, nil
}
