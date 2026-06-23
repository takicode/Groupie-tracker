package config

import "os"

type Config struct{
	Port string
	BaseURL string 
	ApiKey string
	ApiBaseUrl string
}


func Load() *Config{
	port := os.Getenv("PORT")
	if port ==""{
		port = "8080"
	}

	baseurl := os.Getenv("BASE_URL")
	if baseurl ==""{
		baseurl = "https://groupietrackers.herokuapp.com/api"
	}

	apikey := os.Getenv("API_KEY") 
	apiString := os.Getenv("API_BASE_URL")
	if apiString == ""{
		apiString= "https://api.opencagedata.com/geocode/v1/json"
	}

	return &Config{
		Port:port,
		BaseURL:baseurl,
		ApiKey:apikey,
		ApiBaseUrl:apiString,
	}
}