package main

import (
	"log"

	"github.com/DaniYer/GoWeatherNow/internal/app"
)

func main() {
	if err := app.InitApp(); err != nil {
		log.Fatalf("Ошибка запуска приложения: %v", err)
	}
}
