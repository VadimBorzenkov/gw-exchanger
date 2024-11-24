package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser   string
	DBPass   string
	DBName   string
	DBHost   string
	DBPort   string
	GRPCPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}, nil
}
