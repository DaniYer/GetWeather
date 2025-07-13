package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenWeatherKey string
	Port           string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	key := os.Getenv("OPENWEATHER_API_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if key == "" {
		return nil, fmt.Errorf("не найден OPENWEATHER_API_KEY")
	}

	return &Config{
		OpenWeatherKey: key,
		Port:           port,
	}, nil
}
