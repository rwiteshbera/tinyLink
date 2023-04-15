package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SENDER_GMAIL    string
	SENDER_PASSWORD string
	SERVER_HOST     string
	SERVER_PORT     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		SENDER_GMAIL:    os.Getenv("SENDER_GMAIL"),
		SENDER_PASSWORD: os.Getenv("SENDER_PASSWORD"),
		SERVER_HOST:     os.Getenv("SERVER_HOST"),
		SERVER_PORT:     os.Getenv("SERVER_PORT"),
	}

	return config, nil
}
