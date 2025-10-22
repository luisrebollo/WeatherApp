/*
This package handles the http requests
*/
package api 

import(
	"time"
	"net/http"
	"encoding/json"

	"weather_app/internal/models"
	"weather_app/internal/weather"
	"weather_app/internal/db"
)

type Handler struct{
	DB db.DBInterface
}

type apiResponse struct{
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

//Gets the weather information when the method is called
func (h *Handler) GetWeatherHandler(w http.ResponseWriter,r *http.Request){
	var req models.WeatherInput

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.City == ""{
		http.Error(w,"Invalid request", http.StatusInternalServerError)
		return
	}
	
	weatherData, err := weather.GetWeatherInfo(req)
	if err != nil{
		http.Error(w, "Failed to get weather", http.StatusInternalServerError)
		return
	}

	query := &models.WeatherQuery{
		City:        req.City,
		Temperature: weatherData.Current.Temp,
		Description: weatherData.Current.Condition.Text,
		Timestamp:   time.Now(),
		IP:          r.RemoteAddr,
	}

	if err := h.DB.SaveQuery(query); err != nil{
		http.Error(w,"Failed to save query",http.StatusInternalServerError)
		return
	}

	resp := &apiResponse{
		City: req.City,
		Temperature: weatherData.Current.Temp,
		Description: weatherData.Current.Condition.Text,
	}



	WriteJson(w,http.StatusOK,resp)

}
//Function that gets the query history 
func (h *Handler) GetWeatherHistoryHandler(w http.ResponseWriter,r *http.Request){
	history, err := h.DB.GetWeatherHistory()
	if err != nil{
		http.Error(w,"Error fetching history", http.StatusInternalServerError)
		return
	}
	WriteJson(w,http.StatusOK,history)
}