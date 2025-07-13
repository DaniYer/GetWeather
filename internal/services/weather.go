package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DaniYer/GoWeatherNow/internal/models"
)

func FetchWeather(city, apiKey string) (*models.CustomResponse, error) {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d from weather API", resp.StatusCode)
	}

	var data models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	description := data.Weather[0].Description

	return &models.CustomResponse{
		City:          data.Name,
		Temperature:   fmt.Sprintf("%.0f°C", data.Main.Temp),
		Description:   description,
		Advice:        getAdvice(description),
		SimilarCities: getSimilarCities(data.Name),
	}, nil
}

func getAdvice(description string) string {
	switch description {
	case "clear sky":
		return "Don't forget your sunglasses!"
	case "rain":
		return "Grab an umbrella!"
	case "snow":
		return "Bundle up, it’s snowing!"
	default:
		return "Be prepared for the weather!"
	}
}

func getSimilarCities(city string) []string {
	switch city {
	case "Tokyo":
		return []string{"Osaka", "Seoul", "Taipei"}
	case "New York":
		return []string{"Boston", "Philadelphia", "Washington"}
	default:
		return []string{"City A", "City B", "City C"}
	}
}
