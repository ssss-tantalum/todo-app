package core

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Debug bool `env:"DEBUG"`

	HTTPPort  int    `env:"HTTP_PORT"`
	SecretKey string `env:"SECRET_KEY"`

	DBdsn string `env:"DB_DSN"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &Config{
		Debug:     cfg.Debug,
		HTTPPort:  cfg.HTTPPort,
		SecretKey: cfg.SecretKey,
		DBdsn:     cfg.DBdsn,
	}
}
