package services

import (
	"github.com/yuseferi/go_api_test/api/domain/weather_domain"
	"github.com/yuseferi/go_api_test/api/providers/weather_provider"
)

type weatherService struct{}

type weatherServiceInterface interface {
	GetWeather(input weather_domain.WeatherRequest) (*weather_domain.Weather, weather_domain.WeatherErrorInterface)
}

var (
	WeatherService weatherServiceInterface = &weatherService{}
)

func (w *weatherService) GetWeather(input weather_domain.WeatherRequest) (*weather_domain.Weather, weather_domain.WeatherErrorInterface) {
	request := weather_domain.WeatherRequest{
		ApiKey:    input.ApiKey,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}
	response, err := weather_provider.WeatherProvider.GetWeather(request)
	if err != nil {
		return nil, weather_domain.NewWeatherError(err.Code, err.ErrorMessage)
	}
	result := weather_domain.Weather{
		Latitude:  response.Latitude,
		Longitude: response.Longitude,
		TimeZone:  response.TimeZone,
		Currently: weather_domain.CurrentlyInfo{
			Temperature: response.Currently.Temperature,
			Summary:     response.Currently.Summary,
			DewPoint:    response.Currently.DewPoint,
			Pressure:    response.Currently.Pressure,
			Humidity:    response.Currently.Humidity,
		},
	}
	return &result, nil
}
