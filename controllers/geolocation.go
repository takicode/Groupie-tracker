package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/api"
	"net/http"
	"net/url"
	"time"
)

func GeoLocation(loc string) ([]api.GeoResponse, error) {
	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	
	query := url.QueryEscape(loc)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", query)

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

	var geoLocate []api.GeoResponse

	
	err = json.NewDecoder(resp.Body).Decode(&geoLocate)
	if err != nil {
		return nil, err
	}

	return geoLocate, nil
}