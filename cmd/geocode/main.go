package main

import (
	"groupie-tracker/internal/artist"
	"groupie-tracker/internal/config"
	"context"
	"log"
	"time"
)

func main(){
	cfg := config.Load()
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second) 
	defer cancel()

	geoclient := artist.NewGeoClient(cfg.ApiKey, cfg.ApiBaseUrl)
    client := artist.NewClient(cfg.BaseURL)
    relations, err := client.GetRelations(ctx)
    if err != nil {
		log.Fatal(err)
	}

	locations := artist.UniqueLocations(relations)
	coords, err := artist.LoadCoordinates()
	
	var missing []string

	for _, locs := range locations{

		
		if _, ok := coords[locs]; !ok{
			missing = append(missing, locs)
		} 

		geo , err := geoclient.GetCoordinates(ctx, locs)

        if err!= nil{
			log.Printf("failed to geocode %s: %v", locs, err)
			continue
		}

        coords[locs] = geo

	}
	
	err = artist.SaveCoordinates(coords)

	if err != nil{
		log.Fatal()
	}

}