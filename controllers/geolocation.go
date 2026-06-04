package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api"
	"net/http"
	"net/url"
	"time"
	"os"
)

func GeoLocation(loc string) ([]api.GeoResponse, error) {

	apiKey := os.Getenv("GeoCode_apiKey")

	if apiKey == "" {
		return nil, fmt.Errorf("GeoCode_apiKey not set")
	}

	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	
	query := url.QueryEscape(loc)
	address := fmt.Sprintf("https://api.opencagedata.com/geocode/v1/json?q=%s&key=%s&language=en&pretty=1", query,apiKey)

	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return []api.GeoResponse{}, err
	}

	req.Header.Set("User-Agent", "Groupie-tracker/1.0")

	resp, err := httpClient.Do(req)
	if err != nil {
		return []api.GeoResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status %s", resp.Status)
	}

	var result api.OpenCageResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	
	var geoLocate []api.GeoResponse

	for _, r := range result.Results {
		geoLocate = append(geoLocate, api.GeoResponse{
			Lat:         fmt.Sprintf("%f", r.Geometry.Lat),
			Lon:         fmt.Sprintf("%f", r.Geometry.Lng),
			DisplayName: r.Formatted,
		})
	}

   return geoLocate, nil

}