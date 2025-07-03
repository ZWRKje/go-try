package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Can't read .env file")
	}

	value := os.Getenv("KEY")
	if value == "" {
		panic("Can't read key in .env file")
	}

	return &Config{
		Key: value,
	}
}
