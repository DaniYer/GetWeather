package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DaniYer/GetWeather.git/internal/config"
	"github.com/DaniYer/GetWeather.git/internal/services"
	"github.com/go-chi/chi/v5"
)

func WeatherHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		city := chi.URLParam(r, "city")
		if city == "" {
			http.Error(w, "Missing city parameter", http.StatusBadRequest)
			return
		}

		resp, err := services.FetchWeather(city, cfg.OpenWeatherKey)
		if err != nil {
			http.Error(w, "Ошибка при получении погоды: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
