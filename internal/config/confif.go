package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBFile   string
	Password string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		Port:     os.Getenv("TODO_PORT"),
		DBFile:   os.Getenv("TODO_DBFILE"),
		Password: os.Getenv("TODO_PASSWORD"),
	}

	if cfg.Port == "" {
		cfg.Port = "7540"
	}

	if cfg.DBFile == "" {
		cfg.DBFile = "scheduler.db"
	}

	return &cfg, nil
}
