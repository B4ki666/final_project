package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBFile string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := Config{
		Port:   os.Getenv("TODO_PORT"),
		DBFile: os.Getenv("TODO_DBFILE"),
	}

	if cfg.Port == "" {
		cfg.Port = "7540"
	}

	if cfg.DBFile == "" {
		cfg.DBFile = "scheduler.db"
	}

	return &cfg, nil
}
