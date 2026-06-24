package artist

import (
	"sync"
	"context"
	// "log"
)

type Store struct{
	client ArtistClient
    artists []FullArtistInfo
	// geo GeoClient
}

type ArtistClient interface{
	GetArtists(ctx context.Context) ([]Artist, error)
	GetRelations(ctx context.Context)([]Relation, error)
}

// type GeoClient interface{
// 	GetCoordinates(ctx context.Context,location string) (GeoLocation, error)
// }

func NewStore(client ArtistClient) *Store{
	return &Store{
		client:client,
		artists:make([]FullArtistInfo, 0),
		// geo:geo,
	}
}

func (s *Store) Load(ctx context.Context)error {

	var( 
	artists   []Artist
	relations []Relation
	artistErr   error
	relationErr error
	wg sync.WaitGroup
)


	wg.Add(2)

	go func(){
		defer wg.Done()
		artists, artistErr = s.client.GetArtists(ctx)
	}()

	go func(){
		defer wg.Done()
		relations, relationErr = s.client.GetRelations(ctx)  
	}()
			
	wg.Wait()
	
	if artistErr != nil {
		return artistErr
	}
	
	if relationErr != nil {
		return relationErr
	}
			
   result := make([]FullArtistInfo,0,len(artists))
   relMap := make(map[int]Relation, len(relations))
//    globalGeoCache := make(map[string]GeoLocation)

   for _, rel := range relations{
	    relMap[rel.ID] = rel
   }

   for _, artist := range artists{
	rel, ok:=relMap[artist.ID]
	if !ok{
		continue
	}

	// coords := make(map[string]GeoLocation)

	// for location := range rel.DatesLocations{

	// 		if geo,ok:=globalGeoCache[location]; ok{
	// 			coords[location] = geo
	// 			continue
	// 		}

	// 		geo,err:= s.geo.GetCoordinates(ctx, location)
	// 		if err != nil{
	// 			log.Printf("error fetching coordinate", err)
	// 			continue
	// 		}
	// 		coords[location] = geo
	// 		globalGeoCache[location]=geo
	// }

	info := FullArtistInfo{
		Artist:artist,
		DatesLocations:rel.DatesLocations,
		// Coordinates:coords,
	}

     result = append(result, info)

	}
	s.artists = result
    return nil
}

func (s *Store) Artists() []FullArtistInfo{
      result := make([]FullArtistInfo, len(s.artists))
	  copy(result, s.artists)
	  return result
   }