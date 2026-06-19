package config

import "os"

type Config struct{
	Port string
	BaseURL string 
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

	return &Config{
		Port:port,
		BaseURL:baseurl,
	}
}