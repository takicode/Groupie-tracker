package controllers


import (
	"strings"
	"log"
	"groupie-tracker/api"
	"time"
	"os"
	"encoding/json"
	"sync"
)

var (
	CoordinatesMap = make(map[string]api.GeoLocation)
	mu sync.RWMutex
)

func Coordinate(locations []string){
    var wg sync.WaitGroup
	batchSize := 10


	for i:= 0; i < len(locations); i+=batchSize{

		end := i+batchSize

		if end > len(locations){
			end = len(locations)
		}
 

		batch := locations[i:end]

		for _, loc :=range batch{
			wg.Add(1)

			go func(locs string){
				defer wg.Done()

			location:= strings.ReplaceAll(locs,"_"," ")
			location = strings.ReplaceAll(location,"-"," ")
	         

			locationResp,err := GeoLocation(location)

			if err != nil {
				log.Println("unable to collect location:", err)
				return
			}

			if len(locationResp) == 0{
				log.Println("location not found",location)
				return
			}

			geo, err := GeoJson(locationResp)

			if err != nil {
				log.Println("GeoJson failed",err)
				return
			}
             
			mu.Lock()
			CoordinatesMap[locs] = geo
			mu.Unlock()

			log.Printf("geocoded: %s", location)

			}(loc)

		}

        wg.Wait()

		if end < len(locations){
			log.Printf("batch %d-%d done, waiting 2s...", i, end)
			time.Sleep(1500 * time.Millisecond)
		}
	}
}

func GetCoords(loc string) (api.GeoLocation, bool) {
    mu.RLock()
    defer mu.RUnlock()
    geo, ok := CoordinatesMap[loc]
    return geo, ok
}


const cacheFile = "coords_cache.json"

func LoadOrBuildCache(locations []string) error {
    // try to load from file first
    if data, err := os.ReadFile(cacheFile); err == nil {

		temp := make(map[string]api.GeoLocation)

		if err := json.Unmarshal(data, &temp); err == nil {
			mu.Lock()
			CoordinatesMap = temp
			defer mu.Unlock()
			log.Printf("loaded %d coords from cache file", len(CoordinatesMap))
			return nil
		}
    }

    // file doesn't exist — geocode and save
    log.Println("cache file not found — geocoding...")
    Coordinate(locations)

    // save to file so next restart is instant
    mu.RLock()
	data, err := json.Marshal(CoordinatesMap)
	mu.RUnlock()
    if err != nil {
        return err
    }

    return os.WriteFile(cacheFile, data, 0644)
}