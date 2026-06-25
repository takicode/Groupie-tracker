package artist

import (
	"sync"
	"context"
	"log"
)

type Store struct{
	client ArtistClient
    artists []FullArtistInfo
	coord CoordinateMap
}

type ArtistClient interface{
	GetArtists(ctx context.Context) ([]Artist, error)
	GetRelations(ctx context.Context)([]Relation, error)
}

func NewStore(client ArtistClient) *Store{
	return &Store{
		client:client,
		artists:make([]FullArtistInfo, 0),
		coord: make(CoordinateMap),
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

	coords,err:= LoadCoordinates()
	if err != nil {
        return err
    }

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


   for _, rel := range relations{
	    relMap[rel.ID] = rel
   }

   for _, artist := range artists{
		rel, ok:=relMap[artist.ID]
		if !ok{
			continue
		}

		if err != nil{
			log.Printf("error fetching coordinate:%v", err)
			return err
		}

		info := FullArtistInfo{
			Artist:artist,
			DatesLocations:rel.DatesLocations,
		}
     	result = append(result, info)
	}

	s.coord = coords
	s.artists = result
    return nil
}

func (s *Store) Artists() []FullArtistInfo{
      result := make([]FullArtistInfo, len(s.artists))
	  copy(result, s.artists)
	  return result
}

func (s *Store) ArtistID(ID int) (FullArtistInfo, error){
    for _, artist := range s.artists {
		if artist.ID == ID {
            return artist, nil
		}
	}
	return FullArtistInfo{}, ErrArtistNotFound
}

func (s *Store) CoordinatesForLocations(locations []string) (map[string]GeoLocation){
	result := make(map[string]GeoLocation)
	for _, loc := range locations{
		geo, ok := s.coord[loc]
		if !ok{
			continue
		}
		result[loc] = geo 
	}

	return result
}

