package config 

import(
	"os"
)

type Config struct{
	Port       string
	WeatherAPI string
	WeatherURL string
	DBUser     string
	DBPass     string
	DBHost     string
	DBName     string
}

//Parse the parameters stored in the .env file

func GetApiKey()string{
	return os.Getenv("WEATHER_API_KEY")
}

func GetBaseWeather()string{
	return os.Getenv("WEATHER_BASE_URL")
}

func LoadConfig() *Config{
	return &Config{
		Port:       getEnv("PORT","8080"),
		WeatherAPI: getEnv("WEATHER_API_KEY",""),
		WeatherURL: getEnv("WEATHER_BASE_URL",""),
		DBUser:     getEnv("DB_USER","root"),
		DBPass:     getEnv("DB_PASS",""),
		DBHost:     getEnv("DB_HOST","localhost"),
		DBName:     getEnv("DB_NAME","weatherbd"),
	}
}
// If the key is not found, it assings a predetermined value
func getEnv(key, fallback string) string{
	if value := os.Getenv(key); value != ""{
		return value
	}
	return fallback
}