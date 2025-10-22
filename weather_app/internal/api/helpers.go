package api

import(
	"net/http"
	"encoding/json"
)

// Function to encode the data to json format

func WriteJson(w http.ResponseWriter,status int, data interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}