package weather_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yuseferi/go_api_test/api/domain/weather_domain"
	"github.com/yuseferi/go_api_test/api/services"
	"log"
	"net/http"
	"strconv"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	long, _ := strconv.ParseFloat(vars["longitude"], 64)
	lat, _ := strconv.ParseFloat(vars["latitude"], 64)
	log.Println(vars)
	request := weather_domain.WeatherRequest{
		ApiKey:    vars["apiKey"],
		Latitude:  lat,
		Longitude: long,
	}
	result, apiError := services.WeatherService.GetWeather(request)
	if apiError != nil {
		log.Println(apiError)
		w.WriteHeader(apiError.Status())
		w.Header().Set("Content-Type", "application/json")
		return
	}
	jsonResp, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}
