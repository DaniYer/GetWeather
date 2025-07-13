package main

import (
	"fmt"
	"net/http"

	getweatherhandler "github.com/DaniYer/GetWeather.git/internal/handlers/getWeatherHandler"
	"github.com/go-chi/chi/v5"
)

func RouterInit() {
	chiRouter := chi.NewRouter()
	chiRouter.Get("/weather", getweatherhandler.WeatherHandler)
	fmt.Println("Weather application is running...")
	if err := http.ListenAndServe(":8080", chiRouter); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
