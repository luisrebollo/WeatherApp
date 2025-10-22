package api 

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Makes available the routes for the frontend

func NewRouter(handler *Handler) http.Handler{
	r := mux.NewRouter()

	r.HandleFunc("/api/weather", handler.GetWeatherHandler).Methods("POST")
	r.HandleFunc("/api/weather/history", handler.GetWeatherHistoryHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)
	return r
}