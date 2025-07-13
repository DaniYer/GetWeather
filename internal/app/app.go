package app

import (
	"fmt"
	"net/http"

	"github.com/DaniYer/GetWeather.git/internal/config"
	"github.com/DaniYer/GetWeather.git/internal/handlers"
	"github.com/DaniYer/GetWeather.git/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitApp() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("ошибка загрузки конфигурации: %w", err)
	}

	// ⚡️ Инициализация логгера
	logger, err := zap.NewProduction() // или zap.NewDevelopment() для локальной отладки
	if err != nil {
		return fmt.Errorf("не удалось инициализировать логгер: %w", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	middlewares.InitLogger(sugar)

	router := chi.NewRouter()

	// Подключение middleware логирования и gzip
	router.Use(middlewares.WithLogging)
	router.Use(middlewares.GzipHandle)

	router.Get("/weather/{city}", handlers.WeatherHandler(cfg))

	sugar.Infof("Сервер стартует на порту %s", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, router)
}
