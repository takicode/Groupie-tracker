package main

import (
	"github.com/joho/godotenv"
	"groupie-tracker/internal/artist"
	"groupie-tracker/internal/config"
	"context"
	"log"
	"time"
	"sync"
	
)

func main(){
	if err := godotenv.Load(); err != nil {
    log.Println("No .env file, use environmental variable")
  }
	cfg := config.Load()
	ctx,cancel := context.WithTimeout(context.Background(), 30 * time.Second) 
	defer cancel()
    
	geoClient := artist.NewGeoClient(cfg.ApiKey, cfg.ApiBaseUrl)
    client := artist.NewClient(cfg.BaseURL)
    relations, err := client.GetRelations(ctx)
    if err != nil {
		log.Printf("failed to get relations  %v", err)
	}

	locations := artist.UniqueLocations(relations)
	log.Println("loading geo locations from opencage api ")
	coords, err := artist.LoadCoordinates()
	 if err != nil {
		log.Printf("failed to load coordinates:%v", err)
	}

	var missing []string

	for _, locs := range locations{
		if _, ok := coords[locs]; !ok{
			missing = append(missing, locs)
		} 
	}
	
	jobs := make(chan string)
	results := make(chan artist.GeoResult)
	var wg sync.WaitGroup

	
	for i := 0; i < 10; i++{
		wg.Add(1)
		go worker(ctx,geoClient,jobs, results, &wg)
	}
	
	go func(){
		for _,locs := range missing{
			jobs <- locs
		}
		close(jobs)
	}()
   
	go func(){
		wg.Wait()
		close(results)
	}()

  
   for r :=range results{
	   if r.Err != nil {
		log.Printf("error getting coordinate:%v", err)
    	continue
		}
		coords[r.Location]=r.Geo
	}
   
	err = artist.SaveCoordinates(coords)

	if err != nil{
		log.Fatalf("unable to save coordinate:%v", err)
	}
	log.Println("coordinates saved to local file")
}
	

func worker(ctx context.Context,geoClient *artist.GeoClient, jobs <-chan string,results chan<- artist.GeoResult,wg *sync.WaitGroup,
){
	defer wg.Done()
	
	for loc := range jobs{
        geo,err := geoClient.GetCoordinates(ctx,loc)

		results <- artist.GeoResult{
			Location:loc,
            Geo:geo,
			Err:err,
		}
    } 
	
	
}






