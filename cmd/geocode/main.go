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
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second) 
	defer cancel()
    
	geoClient := artist.NewGeoClient(cfg.ApiKey, cfg.ApiBaseUrl)
    client := artist.NewClient(cfg.BaseURL)
    relations, err := client.GetRelations(ctx)
    if err != nil {
		log.Printf("failed to get relations  %v", err)
	}

	locations := artist.UniqueLocations(relations)
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
	results := make(chan GeoResult)
    var wg sync.WaitGroup
	
   for _, loc := range missing{
        jobs <- loc
    } 

   for i = 0;i < 10; i++{
	  wg.Add(1)
      go worker(ctx,geoClient,jobs, results, &wg)
   } 
	wg.Wait()






	

}


func worker(ctx context.Context,geoClient *GeoClient, jobs <-chan string,results chan<- GeoResult,wg *sync.WaitGroup,
){
		defer wg.Done()
		for loc := range jobs{

		}

		close(jobs)
	  
}