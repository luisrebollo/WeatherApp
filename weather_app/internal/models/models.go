/*
This model aims to hold all the structures used on this project
*/
package models

import(
	"time"
)

type WeatherInput struct{
	City     string `json:"city"`
	Language string `json:"language"`
}

type WeatherResponse struct{
	Current Current `json:"current"`
}

type Current struct{
	Temp     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Condition struct{
	Text string `json:"text"`
}

type WeatherQuery struct{
	ID          string    `json:"id"`
	City        string    `json:"city"`
	Temperature float64   `json:"temperature"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"queried_at"`
	IP          string    `json:"ip_address"`

}