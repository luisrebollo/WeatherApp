/*
The package weather is responsible for getting the corresponding 
weather information given a specific location (City)
*/

package weather

import(
	"io"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"weather_app/config"
	"weather_app/internal/models"
)

func GetWeatherInfo(req models.WeatherInput)(models.WeatherResponse,error){
	var response models.WeatherResponse
	var format string
	
	format = "/current.json"

	apiURL := config.GetBaseWeather()
	apiKey := config.GetApiKey()

	urlReq := fmt.Sprintf("%s%s?key=%s&q=%s&lang=%s",
		apiURL, format, apiKey, req.City, req.Language)

	resp, err := http.Get(urlReq)

	if err != nil{
		log.Println("weatherApp - ExtAPI: ",err)
		return response,nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Println("weatherApp - ExtAPI: ",err)
		return response,nil
	}

	err = json.Unmarshal(body,&response)
	if err != nil{
		log.Println("weatherApp - ExtAPI: ",err)
	}
	return response,nil
}

