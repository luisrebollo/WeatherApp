/*
Package to orchestrate the handlers and make the frontend available
*/
package main

import (
	"log"
	"net/http"
    "github.com/joho/godotenv"
	"weather_app/internal/api"
	"weather_app/internal/db"
	"weather_app/config"
)
//Gets the  environment variables
func initEnv(){
	err := godotenv.Load("/home/ubuntu/weather_app/config/.env")
	if err != nil{
		log.Println("weather_app - config: ",err)
	}
}

func main(){
	initEnv()
	//Configuration parameters
	cfg := config.LoadConfig()
	//Database object
	database, err := db.New(cfg)
	if err != nil{
		log.Fatal(err)
	}
	//Close database object at the end of the execution
	defer database.Close()
	//Handler object with database object as a parameter
	handler := &api.Handler{
		DB: database,
	}
	//Makes the rest api available
	router := api.NewRouter(handler)

	addr := ":" + cfg.Port	
	log.Println("Server listening on ", addr)
	http.ListenAndServe(addr,router)
}