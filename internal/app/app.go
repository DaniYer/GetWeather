package app

import (
	"fmt"
	"net/http"

	"github.com/DaniYer/GoWeatherNow/internal/config"
	"github.com/DaniYer/GoWeatherNow/internal/handlers"
	"github.com/DaniYer/GoWeatherNow/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func InitApp() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("ошибка загрузки конфигурации: %w", err)
	}

	router := chi.NewRouter()
	router.Use(middlewares.GzipHandle)

	router.Get("/weather/{city}", handlers.WeatherHandler(cfg))

	fmt.Printf("🚀 Сервер стартует на порту %s\n", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, router)
}
