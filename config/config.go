package config

import (
	"fmt"
	"os"

	"github.com/0x6969766f/attendant/database"
	"github.com/joho/godotenv"
)

type Config struct {
	PSQL   database.Config
	Server struct {
		Address string
	}
}

func Load() (Config, error) {
	var config Config

	err := godotenv.Load()
	if err != nil {
		return config, nil
	}

	config.PSQL = database.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
		SSLMode:  os.Getenv("PG_SSLMODE"),
	}
	if config.PSQL.Host == "" && config.PSQL.Port == "" {
		return config, fmt.Errorf("missing database config")
	}

	config.Server.Address = os.Getenv("SERVER_ADDRESS")

	return config, nil
}
