package models

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type CustomResponse struct {
	City          string   `json:"city"`
	Temperature   string   `json:"temperature"`
	Description   string   `json:"description"`
	Advice        string   `json:"advice"`
	SimilarCities []string `json:"similarCities"`
}
