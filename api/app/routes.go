package app

import (
	"github.com/gorilla/mux"
	"github.com/yuseferi/go_api_test/api/controllers/weather_controller"
	"net/http"
)

func routes() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/weather/{apiKey}/{latitude}/{longitude}", weather_controller.GetWeather).Methods(http.MethodGet)
	r.HandleFunc("/test", weather_controller.GetWeather).Methods(http.MethodGet)
	http.Handle("/", r)
	return
}
