package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host      string `envconfig:"HOST"`
	Port      string `envconfig:"PORT"`
	DBUser    string `envconfig:"DB_USER"`
	DBPass    string `envconfig:"DB_PASS"`
	DBHost    string `envconfig:"DB_HOST"`
	DBPort    string `envconfig:"DB_PORT"`
	DBName    string `envconfig:"DB_NAME"`
	Salt      string `envconfig:"SALT"`
	SecretKey string `envconfig:"SECRET_KEY"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w ", err)
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("Error loading config: %w ", err)
	}

	return cfg, nil
}
