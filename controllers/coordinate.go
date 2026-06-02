package controllers


import (
	"strings"
	"log"
	"groupie-tracker/api"
	"time"
	"os"
	"encoding/json"
)

var CoordinatesMap = make(map[string]api.GeoLocation)

func Coordinate(locations []string){
	for _, loc := range locations{
		location:= strings.ReplaceAll(loc,"_"," ")
		location = strings.ReplaceAll(location,"-"," ")

		locationResp,err := GeoLocation(location)
			if err != nil {
			log.Println("unable to collect location:", err)
			continue
		}
		if len(locationResp) == 0{
			 log.Println("location not found")
			continue
		}

		geo, err := GeoJson(locationResp)
			if err != nil {
				continue
			}

    	CoordinatesMap[loc] = geo

		 
        time.Sleep(1 * time.Second)
	}
}


const cacheFile = "coords_cache.json"

func LoadOrBuildCache(locations []string) error {
    // try to load from file first
    if data, err := os.ReadFile(cacheFile); err == nil {
        if err := json.Unmarshal(data, &CoordinatesMap); err == nil {
            log.Printf("loaded %d coords from cache file", len(CoordinatesMap))
            return nil
        }
    }

    // file doesn't exist — geocode and save
    log.Println("cache file not found — geocoding...")
    Coordinate(locations)

    // save to file so next restart is instant
    data, err := json.Marshal(CoordinatesMap)
    if err != nil {
        return err
    }

    return os.WriteFile(cacheFile, data, 0644)
}